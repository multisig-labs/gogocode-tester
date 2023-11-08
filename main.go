package main

import (
	"strings"

	"flag"
	"fmt"
	"os"

	"runtime/debug"

	"github.com/jxskiss/mcli"
	"github.com/multisig-labs/blockchain-tester/internal"
	"github.com/multisig-labs/blockchain-tester/pkg/version"
)

func main() {
	defer handlePanic()
	mcli.Add("test", testCmd, "Run the tester against a binary")
	mcli.Add("version", versionCmd, "Display version")
	mcli.AddHelp()
	mcli.AddCompletion()
	mcli.Run()
}

func testCmd() {
	args := struct {
		Stage      string `cli:"--stage, Stage to test"`
		Executable string `cli:"--exe, Executable to test"`
	}{}
	mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))
	handleError(internal.RunTestCases(args.Executable, args.Stage))
}

func versionCmd() {
	fmt.Println("Build Date:", version.BuildDate)
	fmt.Println("Git Commit:", version.GitCommit)
	fmt.Println("Version:", version.Version)
	fmt.Println("Go Version:", version.GoVersion)
	fmt.Println("OS / Arch:", version.OsArch)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handlePanic() {
	if panicPayload := recover(); panicPayload != nil {
		stack := string(debug.Stack())
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintln(os.Stderr, "            Fatal error. Sorry! You found a bug.")
		fmt.Fprintln(os.Stderr, "    Please copy all of this info into an issue at")
		fmt.Fprintln(os.Stderr, "     https://github.com/multisig-labs")
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintf(os.Stderr, "Version:           %s\n", version.Version)
		fmt.Fprintf(os.Stderr, "Build Date:        %s\n", version.BuildDate)
		fmt.Fprintf(os.Stderr, "Git Commit:        %s\n", version.GitCommit)
		fmt.Fprintf(os.Stderr, "Go Version:        %s\n", version.GoVersion)
		fmt.Fprintf(os.Stderr, "OS / Arch:         %s\n", version.OsArch)
		fmt.Fprintf(os.Stderr, "Panic:             %s\n\n", panicPayload)
		fmt.Fprintln(os.Stderr, stack)
		os.Exit(1)
	}
}

func envMap() map[string]string {
	result := make(map[string]string)
	for _, keyVal := range os.Environ() {
		split := strings.SplitN(keyVal, "=", 2)
		key, val := split[0], split[1]
		result[key] = val
	}

	return result
}
