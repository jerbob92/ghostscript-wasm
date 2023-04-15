package main

import (
	"context"
	"crypto/rand"
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

        "github.com/jerbob92/ghostscript-wasm/imports"

        "github.com/tetratelabs/wazero"
        "github.com/tetratelabs/wazero/experimental"
        "github.com/tetratelabs/wazero/experimental/logging"
        "github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed wasm/gs.wasm
var gsWasm []byte

func main() {
	ctx := context.WithValue(context.Background(), experimental.FunctionListenerFactoryKey{}, logging.NewLoggingListenerFactory(os.Stdout))
	ctx = context.Background() // Comment this line to get debug information.

	runtimeConfig := wazero.NewRuntimeConfigInterpreter()
	//cache, err := wazero.NewCompilationCacheWithDir(".wazero-cache")
	//if err == nil {
	//		runtimeConfig = runtimeConfig.WithCompilationCache(cache)
	//	}
	wazeroRuntime := wazero.NewRuntimeWithConfig(ctx, runtimeConfig)

	defer wazeroRuntime.Close(ctx)

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, wazeroRuntime); err != nil {
		log.Fatal(err)
	}

	compiledModule, err := wazeroRuntime.CompileModule(ctx, gsWasm)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := imports.Instantiate(ctx, wazeroRuntime); err != nil {
		log.Fatal(err)
	}

	fsConfig := wazero.NewFSConfig()

	// On Windows we mount the volume of the current working directory as
	// root. On Linux we mount / as root.
	if runtime.GOOS == "windows" {
		cwdDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		volumeName := filepath.VolumeName(cwdDir)
		if volumeName != "" {
			fsConfig = fsConfig.WithDirMount(fmt.Sprintf("%s\\", volumeName), "/")
		}
	} else {
		fsConfig = fsConfig.WithDirMount("/", "/")
	}

	args := []string{"gs"}
	args = append(args, os.Args[1:]...)

	moduleConfig := wazero.NewModuleConfig().
		WithStartFunctions("_start").
		WithStdout(os.Stdout).
		WithStderr(os.Stderr).
		WithRandSource(rand.Reader).
		WithFSConfig(fsConfig).
		WithName("").
		WithArgs(args...)

	_, err = wazeroRuntime.InstantiateModule(ctx, compiledModule, moduleConfig)
	if err != nil {
		log.Fatal(err)
	}
}
