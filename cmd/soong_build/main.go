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

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/blueprint/bootstrap"

	"android/soong/android"
)

var (
	docFile         string
	bazelOverlayDir string
)

func init() {
	flag.StringVar(&docFile, "soong_docs", "", "build documentation file to output")
	flag.StringVar(&bazelOverlayDir, "bazel_overlay_dir", "", "path to the bazel overlay directory")
}

func newNameResolver(config android.Config) *android.NameResolver {
	namespacePathsToExport := make(map[string]bool)

	for _, namespaceName := range config.ExportedNamespaces() {
		namespacePathsToExport[namespaceName] = true
	}

	namespacePathsToExport["."] = true // always export the root namespace

	exportFilter := func(namespace *android.Namespace) bool {
		return namespacePathsToExport[namespace.Path]
	}

	return android.NewNameResolver(exportFilter)
}

func main() {
	android.ReexecWithDelveMaybe()
	flag.Parse()

	// The top-level Blueprints file is passed as the first argument.
	srcDir := filepath.Dir(flag.Arg(0))

	ctx := android.NewContext()
	ctx.Register()

	configuration, err := android.NewConfig(srcDir, bootstrap.BuildDir, bootstrap.ModuleListFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	if !shouldPrepareBuildActions() {
		configuration.SetStopBefore(bootstrap.StopBeforePrepareBuildActions)
	}

	ctx.SetNameInterface(newNameResolver(configuration))

	ctx.SetAllowMissingDependencies(configuration.AllowMissingDependencies())

	extraNinjaDeps := []string{configuration.ConfigFileName, configuration.ProductVariablesFileName}

	// Read the SOONG_DELVE again through configuration so that there is a dependency on the environment variable
	// and soong_build will rerun when it is set for the first time.
	if listen := configuration.Getenv("SOONG_DELVE"); listen != "" {
		// Add a non-existent file to the dependencies so that soong_build will rerun when the debugger is
		// enabled even if it completed successfully.
		extraNinjaDeps = append(extraNinjaDeps, filepath.Join(configuration.BuildDir(), "always_rerun_for_delve"))
	}

	bootstrap.Main(ctx.Context, configuration, extraNinjaDeps...)

	if bazelOverlayDir != "" {
		if err := createBazelOverlay(ctx, bazelOverlayDir); err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
	}

	if docFile != "" {
		if err := writeDocs(ctx, docFile); err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
	}

	// TODO(ccross): make this a command line argument.  Requires plumbing through blueprint
	//  to affect the command line of the primary builder.
	if shouldPrepareBuildActions() {
		metricsFile := filepath.Join(bootstrap.BuildDir, "soong_build_metrics.pb")
		err = android.WriteMetrics(configuration, metricsFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing soong_build metrics %s: %s", metricsFile, err)
			os.Exit(1)
		}
	}
}

func shouldPrepareBuildActions() bool {
	// If we're writing soong_docs or bazel_overlay, don't write build.ninja or
	// collect metrics.
	return docFile == "" && bazelOverlayDir == ""
}
