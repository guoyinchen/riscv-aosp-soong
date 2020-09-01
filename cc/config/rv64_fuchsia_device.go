// Copyright 2018 Google Inc. All rights reserved.
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
	"android/soong/android"
)

var fuchsiaRv64SysRoot string = "prebuilts/fuchsia_sdk/arch/rv64/sysroot"
var fuchsiaRv64PrebuiltLibsRoot string = "fuchsia/prebuilt_libs/"

type toolchainFuchsiaRv64 struct {
	toolchain64Bit
	toolchainFuchsia
}

func (t *toolchainFuchsiaRv64) Name() string {
	return "rv64"
}

func (t *toolchainFuchsiaRv64) GccRoot() string {
	return "${config.Rv64GccRoot}"
}

func (t *toolchainFuchsiaRv64) GccTriple() string {
	return "rv64-linux-android"
}

func (t *toolchainFuchsiaRv64) GccVersion() string {
	return rv64GccVersion
}

func (t *toolchainFuchsiaRv64) Cflags() string {
	return ""
}

func (t *toolchainFuchsiaRv64) Cppflags() string {
	return ""
}

func (t *toolchainFuchsiaRv64) IncludeFlags() string {
	return ""
}

func (t *toolchainFuchsiaRv64) ClangTriple() string {
	return "rv64-fuchsia-android"
}

func (t *toolchainFuchsiaRv64) ClangCppflags() string {
	return "-Wno-error=deprecated-declarations"
}

func (t *toolchainFuchsiaRv64) ClangLdflags() string {
	return "--target=rv64-fuchsia --sysroot=" + fuchsiaRv64SysRoot + " -L" + fuchsiaRv64PrebuiltLibsRoot + "/rv64-fuchsia/lib " + "-Lprebuilts/fuchsia_sdk/arch/rv64/dist/"
}

func (t *toolchainFuchsiaRv64) ClangLldflags() string {
	return "--target=rv64-fuchsia --sysroot=" + fuchsiaRv64SysRoot + " -L" + fuchsiaRv64PrebuiltLibsRoot + "/rv64-fuchsia/lib " + "-Lprebuilts/fuchsia_sdk/arch/rv64/dist/"
}

func (t *toolchainFuchsiaRv64) ClangCflags() string {
	return "--target=rv64-fuchsia --sysroot=" + fuchsiaRv64SysRoot + " -I" + fuchsiaRv64SysRoot + "/include"
}

func (t *toolchainFuchsiaRv64) Bionic() bool {
	return false
}

func (t *toolchainFuchsiaRv64) ToolchainClangCflags() string {
	return "-march=rv64i"
}

var toolchainRv64FuchsiaSingleton Toolchain = &toolchainFuchsiaRv64{}

func rv64FuchsiaToolchainFactory(arch android.Arch) Toolchain {
	return toolchainRv64FuchsiaSingleton
}

func init() {
	registerToolchainFactory(android.Fuchsia, android.Rv64, rv64FuchsiaToolchainFactory)
}
