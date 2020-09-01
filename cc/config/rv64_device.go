// Copyright 2015 Google Inc. All rights reserved.
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

package config

import (
	"fmt"
	"strings"

	"android/soong/android"
)

var (
	rv64Cflags = []string{
		// Help catch common 32/64-bit errors.
		"-Werror=implicit-function-declaration",
	}

	rv64ArchVariantCflags = map[string][]string{
		"rv64i": []string{
			"-march=rv64i",
		},
		"rv64im": []string{
			"-march=rv64im",
		},
	}

	rv64Ldflags = []string{
		"-Wl,--hash-style=gnu",
		"-fuse-ld=gold",
		"-Wl,--icf=safe",
	}

	rv64Lldflags = append(ClangFilterUnknownLldflags(rv64Ldflags),
		"-Wl,-z,max-page-size=4096")

	rv64Cppflags = []string{}
)

const (
	rv64GccVersion = "4.9"
)

func init() {
	pctx.StaticVariable("rv64GccVersion", rv64GccVersion)

	pctx.SourcePathVariable("Rv64GccRoot",
		"prebuilts/gcc/${HostPrebuiltTag}/rv64/rv64-linux-android-${rv64GccVersion}")

	pctx.StaticVariable("Rv64Ldflags", strings.Join(rv64Ldflags, " "))
	pctx.StaticVariable("Rv64Lldflags", strings.Join(rv64Lldflags, " "))
	pctx.StaticVariable("Rv64IncludeFlags", bionicHeaders("rv64"))

	pctx.StaticVariable("Rv64ClangCflags", strings.Join(ClangFilterUnknownCflags(rv64Cflags), " "))
	pctx.StaticVariable("Rv64ClangLdflags", strings.Join(ClangFilterUnknownCflags(rv64Ldflags), " "))
	pctx.StaticVariable("Rv64ClangLldflags", strings.Join(ClangFilterUnknownCflags(rv64Lldflags), " "))
	pctx.StaticVariable("Rv64ClangCppflags", strings.Join(ClangFilterUnknownCflags(rv64Cppflags), " "))

	pctx.StaticVariable("Rv64ClangRv64ICflags", strings.Join(rv64ArchVariantCflags["rv64i"], " "))
	pctx.StaticVariable("Rv64ClangRv64IMCflags", strings.Join(rv64ArchVariantCflags["rv64im"], " "))

}

var (
	rv64ClangArchVariantCflagsVar = map[string]string{
		"rv64i":  "${config.Rv64ClangRv64ICflags}",
		"rv64im": "${config.Rv64ClangRv64IMCflags}",
	}

	rv64ClangCpuVariantCflagsVar = map[string]string{
		"":           "",
	}
)

type toolchainRv64 struct {
	toolchain64Bit

	ldflags              string
	lldflags             string
	toolchainClangCflags string
}

func (t *toolchainRv64) Name() string {
	return "rv64"
}

func (t *toolchainRv64) GccRoot() string {
	return "${config.Rv64GccRoot}"
}

func (t *toolchainRv64) GccTriple() string {
	return "rv64-linux-android"
}

func (t *toolchainRv64) GccVersion() string {
	return rv64GccVersion
}

func (t *toolchainRv64) IncludeFlags() string {
	return "${config.Rv64IncludeFlags}"
}

func (t *toolchainRv64) ClangTriple() string {
	return t.GccTriple()
}

func (t *toolchainRv64) ClangCflags() string {
	return "${config.Rv64ClangCflags}"
}

func (t *toolchainRv64) ClangCppflags() string {
	return "${config.Rv64ClangCppflags}"
}

func (t *toolchainRv64) ClangLdflags() string {
	return t.ldflags
}

func (t *toolchainRv64) ClangLldflags() string {
	return t.lldflags
}

func (t *toolchainRv64) ToolchainClangCflags() string {
	return t.toolchainClangCflags
}

func (toolchainRv64) LibclangRuntimeLibraryArch() string {
	return "rv64"
}

func rv64ToolchainFactory(arch android.Arch) Toolchain {
	switch arch.ArchVariant {
	case "rv64i":
	case "rv64im":
		// Nothing extra for rv64i/rv64im
	default:
		panic(fmt.Sprintf("Unknown RISC-V architecture version: %q", arch.ArchVariant))
	}

	toolchainClangCflags := []string{rv64ClangArchVariantCflagsVar[arch.ArchVariant]}
	toolchainClangCflags = append(toolchainClangCflags,
		variantOrDefault(rv64ClangCpuVariantCflagsVar, arch.CpuVariant))

	var extraLdflags string

	return &toolchainRv64{
		ldflags: strings.Join([]string{
			"${config.Rv64Ldflags}",
			extraLdflags,
		}, " "),
		lldflags: strings.Join([]string{
			"${config.Rv64Lldflags}",
			extraLdflags,
		}, " "),
		toolchainClangCflags: strings.Join(toolchainClangCflags, " "),
	}
}

func init() {
	registerToolchainFactory(android.Android, android.Rv64, rv64ToolchainFactory)
}
