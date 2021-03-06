// Copyright 2016 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cc

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/google/blueprint"

	"android/soong/android"
)

func init() {
	pctx.HostBinToolVariable("ndkStubGenerator", "ndkstubgen")
	pctx.HostBinToolVariable("ndk_api_coverage_parser", "ndk_api_coverage_parser")
}

var (
	genStubSrc = pctx.AndroidStaticRule("genStubSrc",
		blueprint.RuleParams{
			Command: "$ndkStubGenerator --arch $arch --api $apiLevel " +
				"--api-map $apiMap $flags $in $out",
			CommandDeps: []string{"$ndkStubGenerator"},
		}, "arch", "apiLevel", "apiMap", "flags")

	parseNdkApiRule = pctx.AndroidStaticRule("parseNdkApiRule",
		blueprint.RuleParams{
			Command:     "$ndk_api_coverage_parser $in $out --api-map $apiMap",
			CommandDeps: []string{"$ndk_api_coverage_parser"},
		}, "apiMap")

	ndkLibrarySuffix = ".ndk"

	// Added as a variation dependency via depsMutator.
	ndkKnownLibs = []string{}
	// protects ndkKnownLibs writes during parallel BeginMutator.
	ndkKnownLibsLock sync.Mutex
)

// Creates a stub shared library based on the provided version file.
//
// Example:
//
// ndk_library {
//     name: "libfoo",
//     symbol_file: "libfoo.map.txt",
//     first_version: "9",
// }
//
type libraryProperties struct {
	// Relative path to the symbol map.
	// An example file can be seen here: TODO(danalbert): Make an example.
	Symbol_file *string

	// The first API level a library was available. A library will be generated
	// for every API level beginning with this one.
	First_version *string

	// The first API level that library should have the version script applied.
	// This defaults to the value of first_version, and should almost never be
	// used. This is only needed to work around platform bugs like
	// https://github.com/android-ndk/ndk/issues/265.
	Unversioned_until *string

	// Private property for use by the mutator that splits per-API level. Can be
	// one of <number:sdk_version> or <codename> or "current" passed to
	// "ndkstubgen.py" as it is
	ApiLevel string `blueprint:"mutated"`

	// True if this API is not yet ready to be shipped in the NDK. It will be
	// available in the platform for testing, but will be excluded from the
	// sysroot provided to the NDK proper.
	Draft bool
}

type stubDecorator struct {
	*libraryDecorator

	properties libraryProperties

	versionScriptPath     android.ModuleGenPath
	parsedCoverageXmlPath android.ModuleOutPath
	installPath           android.Path
}

// OMG GO
func intMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func normalizeNdkApiLevel(ctx android.BaseModuleContext, apiLevel string,
	arch android.Arch) (string, error) {

	if apiLevel == "" {
		panic("empty apiLevel not allowed")
	}

	if apiLevel == "current" {
		return apiLevel, nil
	}

	minVersion := ctx.Config().MinSupportedSdkVersion()
	firstArchVersions := map[android.ArchType]int{
		android.Arm:    minVersion,
		android.Arm64:  21,
		android.X86:    minVersion,
		android.X86_64: 21,
	}

	firstArchVersion, ok := firstArchVersions[arch.ArchType]
	if !ok {
		panic(fmt.Errorf("Arch %q not found in firstArchVersions", arch.ArchType))
	}

	if apiLevel == "minimum" {
		return strconv.Itoa(firstArchVersion), nil
	}

	// If the NDK drops support for a platform version, we don't want to have to
	// fix up every module that was using it as its SDK version. Clip to the
	// supported version here instead.
	version, err := strconv.Atoi(apiLevel)
	if err != nil {
		// Non-integer API levels are codenames.
		return apiLevel, nil
	}
	version = intMax(version, minVersion)

	return strconv.Itoa(intMax(version, firstArchVersion)), nil
}

func getFirstGeneratedVersion(firstSupportedVersion string, platformVersion int) (int, error) {
	if firstSupportedVersion == "current" {
		return platformVersion + 1, nil
	}

	return strconv.Atoi(firstSupportedVersion)
}

func shouldUseVersionScript(ctx android.BaseModuleContext, stub *stubDecorator) (bool, error) {
	// unversioned_until is normally empty, in which case we should use the version script.
	if String(stub.properties.Unversioned_until) == "" {
		return true, nil
	}

	if String(stub.properties.Unversioned_until) == "current" {
		if stub.properties.ApiLevel == "current" {
			return true, nil
		} else {
			return false, nil
		}
	}

	if stub.properties.ApiLevel == "current" {
		return true, nil
	}

	unversionedUntil, err := android.ApiStrToNum(ctx, String(stub.properties.Unversioned_until))
	if err != nil {
		return true, err
	}

	version, err := android.ApiStrToNum(ctx, stub.properties.ApiLevel)
	if err != nil {
		return true, err
	}

	return version >= unversionedUntil, nil
}

func generatePerApiVariants(ctx android.BottomUpMutatorContext, m *Module,
	propName string, propValue string, perSplit func(*Module, string)) {
	platformVersion := ctx.Config().PlatformSdkVersionInt()

	firstSupportedVersion, err := normalizeNdkApiLevel(ctx, propValue,
		ctx.Arch())
	if err != nil {
		ctx.PropertyErrorf(propName, err.Error())
	}

	firstGenVersion, err := getFirstGeneratedVersion(firstSupportedVersion,
		platformVersion)
	if err != nil {
		// In theory this is impossible because we've already run this through
		// normalizeNdkApiLevel above.
		ctx.PropertyErrorf(propName, err.Error())
	}

	var versionStrs []string
	for version := firstGenVersion; version <= platformVersion; version++ {
		versionStrs = append(versionStrs, strconv.Itoa(version))
	}
	versionStrs = append(versionStrs, ctx.Config().PlatformVersionActiveCodenames()...)
	versionStrs = append(versionStrs, "current")

	modules := ctx.CreateVariations(versionStrs...)
	for i, module := range modules {
		perSplit(module.(*Module), versionStrs[i])
	}
}

func NdkApiMutator(ctx android.BottomUpMutatorContext) {
	if m, ok := ctx.Module().(*Module); ok {
		if m.Enabled() {
			if compiler, ok := m.compiler.(*stubDecorator); ok {
				if ctx.Os() != android.Android {
					// These modules are always android.DeviceEnabled only, but
					// those include Fuchsia devices, which we don't support.
					ctx.Module().Disable()
					return
				}
				generatePerApiVariants(ctx, m, "first_version",
					String(compiler.properties.First_version),
					func(m *Module, version string) {
						m.compiler.(*stubDecorator).properties.ApiLevel =
							version
					})
			} else if m.SplitPerApiLevel() && m.IsSdkVariant() {
				if ctx.Os() != android.Android {
					return
				}
				generatePerApiVariants(ctx, m, "min_sdk_version",
					m.MinSdkVersion(), func(m *Module, version string) {
						m.Properties.Sdk_version = &version
					})
			}
		}
	}
}

func (c *stubDecorator) compilerInit(ctx BaseModuleContext) {
	c.baseCompiler.compilerInit(ctx)

	name := ctx.baseModuleName()
	if strings.HasSuffix(name, ndkLibrarySuffix) {
		ctx.PropertyErrorf("name", "Do not append %q manually, just use the base name", ndkLibrarySuffix)
	}

	ndkKnownLibsLock.Lock()
	defer ndkKnownLibsLock.Unlock()
	for _, lib := range ndkKnownLibs {
		if lib == name {
			return
		}
	}
	ndkKnownLibs = append(ndkKnownLibs, name)
}

func addStubLibraryCompilerFlags(flags Flags) Flags {
	flags.Global.CFlags = append(flags.Global.CFlags,
		// We're knowingly doing some otherwise unsightly things with builtin
		// functions here. We're just generating stub libraries, so ignore it.
		"-Wno-incompatible-library-redeclaration",
		"-Wno-incomplete-setjmp-declaration",
		"-Wno-builtin-requires-header",
		"-Wno-invalid-noreturn",
		"-Wall",
		"-Werror",
		// These libraries aren't actually used. Don't worry about unwinding
		// (avoids the need to link an unwinder into a fake library).
		"-fno-unwind-tables",
	)
	// All symbols in the stubs library should be visible.
	if inList("-fvisibility=hidden", flags.Local.CFlags) {
		flags.Local.CFlags = append(flags.Local.CFlags, "-fvisibility=default")
	}
	return flags
}

func (stub *stubDecorator) compilerFlags(ctx ModuleContext, flags Flags, deps PathDeps) Flags {
	flags = stub.baseCompiler.compilerFlags(ctx, flags, deps)
	return addStubLibraryCompilerFlags(flags)
}

func compileStubLibrary(ctx ModuleContext, flags Flags, symbolFile, apiLevel, genstubFlags string) (Objects, android.ModuleGenPath) {
	arch := ctx.Arch().ArchType.String()

	stubSrcPath := android.PathForModuleGen(ctx, "stub.c")
	versionScriptPath := android.PathForModuleGen(ctx, "stub.map")
	symbolFilePath := android.PathForModuleSrc(ctx, symbolFile)
	apiLevelsJson := android.GetApiLevelsJson(ctx)
	ctx.Build(pctx, android.BuildParams{
		Rule:        genStubSrc,
		Description: "generate stubs " + symbolFilePath.Rel(),
		Outputs:     []android.WritablePath{stubSrcPath, versionScriptPath},
		Input:       symbolFilePath,
		Implicits:   []android.Path{apiLevelsJson},
		Args: map[string]string{
			"arch":     arch,
			"apiLevel": apiLevel,
			"apiMap":   apiLevelsJson.String(),
			"flags":    genstubFlags,
		},
	})

	subdir := ""
	srcs := []android.Path{stubSrcPath}
	return compileObjs(ctx, flagsToBuilderFlags(flags), subdir, srcs, nil, nil), versionScriptPath
}

func parseSymbolFileForCoverage(ctx ModuleContext, symbolFile string) android.ModuleOutPath {
	apiLevelsJson := android.GetApiLevelsJson(ctx)
	symbolFilePath := android.PathForModuleSrc(ctx, symbolFile)
	outputFileName := strings.Split(symbolFilePath.Base(), ".")[0]
	parsedApiCoveragePath := android.PathForModuleOut(ctx, outputFileName+".xml")
	ctx.Build(pctx, android.BuildParams{
		Rule:        parseNdkApiRule,
		Description: "parse ndk api symbol file for api coverage: " + symbolFilePath.Rel(),
		Outputs:     []android.WritablePath{parsedApiCoveragePath},
		Input:       symbolFilePath,
		Implicits:   []android.Path{apiLevelsJson},
		Args: map[string]string{
			"apiMap": apiLevelsJson.String(),
		},
	})
	return parsedApiCoveragePath
}

func (c *stubDecorator) compile(ctx ModuleContext, flags Flags, deps PathDeps) Objects {
	if !strings.HasSuffix(String(c.properties.Symbol_file), ".map.txt") {
		ctx.PropertyErrorf("symbol_file", "must end with .map.txt")
	}

	symbolFile := String(c.properties.Symbol_file)
	objs, versionScript := compileStubLibrary(ctx, flags, symbolFile,
		c.properties.ApiLevel, "")
	c.versionScriptPath = versionScript
	if c.properties.ApiLevel == "current" && ctx.PrimaryArch() {
		c.parsedCoverageXmlPath = parseSymbolFileForCoverage(ctx, symbolFile)
	}
	return objs
}

func (linker *stubDecorator) linkerDeps(ctx DepsContext, deps Deps) Deps {
	return Deps{}
}

func (linker *stubDecorator) Name(name string) string {
	return name + ndkLibrarySuffix
}

func (stub *stubDecorator) linkerFlags(ctx ModuleContext, flags Flags) Flags {
	stub.libraryDecorator.libName = ctx.baseModuleName()
	return stub.libraryDecorator.linkerFlags(ctx, flags)
}

func (stub *stubDecorator) link(ctx ModuleContext, flags Flags, deps PathDeps,
	objs Objects) android.Path {

	useVersionScript, err := shouldUseVersionScript(ctx, stub)
	if err != nil {
		ctx.ModuleErrorf(err.Error())
	}

	if useVersionScript {
		linkerScriptFlag := "-Wl,--version-script," + stub.versionScriptPath.String()
		flags.Local.LdFlags = append(flags.Local.LdFlags, linkerScriptFlag)
		flags.LdFlagsDeps = append(flags.LdFlagsDeps, stub.versionScriptPath)
	}

	return stub.libraryDecorator.link(ctx, flags, deps, objs)
}

func (stub *stubDecorator) nativeCoverage() bool {
	return false
}

func (stub *stubDecorator) install(ctx ModuleContext, path android.Path) {
	arch := ctx.Target().Arch.ArchType.Name
	apiLevel := stub.properties.ApiLevel

	// arm64 isn't actually a multilib toolchain, so unlike the other LP64
	// architectures it's just installed to lib.
	libDir := "lib"
	if ctx.toolchain().Is64Bit() && arch != "arm64" {
		libDir = "lib64"
	}

	installDir := getNdkInstallBase(ctx).Join(ctx, fmt.Sprintf(
		"platforms/android-%s/arch-%s/usr/%s", apiLevel, arch, libDir))
	stub.installPath = ctx.InstallFile(installDir, path.Base(), path)
}

func newStubLibrary() *Module {
	module, library := NewLibrary(android.DeviceSupported)
	library.BuildOnlyShared()
	module.stl = nil
	module.sanitize = nil
	library.disableStripping()

	stub := &stubDecorator{
		libraryDecorator: library,
	}
	module.compiler = stub
	module.linker = stub
	module.installer = stub

	module.Properties.AlwaysSdk = true
	module.Properties.Sdk_version = StringPtr("current")

	module.AddProperties(&stub.properties, &library.MutatedProperties)

	return module
}

// ndk_library creates a library that exposes a stub implementation of functions
// and variables for use at build time only.
func NdkLibraryFactory() android.Module {
	module := newStubLibrary()
	android.InitAndroidArchModule(module, android.DeviceSupported, android.MultilibBoth)
	module.ModuleBase.EnableNativeBridgeSupportByDefault()
	return module
}
