package main

import (
	"flag"
	"fmt"
	"os"

	"runtime/debug"

	"github.com/jxskiss/mcli"
	"github.com/multisig-labs/gogocode-tester-sample/pkg/testcase_sample"
	"github.com/multisig-labs/gogocode-tester-sample/pkg/version"
)

func main() {
	defer handlePanic()
	mcli.Add("sample", testSampleCmd, "Run the tester for course Sample against a binary. Returns slug of last successful test case")
	mcli.Add("version", versionCmd, "Display version")
	mcli.Run()
}

func testSampleCmd() {
	args := struct {
		Executable string `cli:"--exe, Executable to test" default:"./run.sh"`
	}{}
	mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))

	// Ignore error (its already been logged) and return last successful stage slug
	slug, _ := testcase_sample.RunTestCases(args.Executable)
	fmt.Print(slug)
}

func versionCmd() {
	fmt.Println("Build Date:", version.BuildDate)
	fmt.Println("Git Commit:", version.GitCommit)
	fmt.Println("Version:", version.Version)
	fmt.Println("Go Version:", version.GoVersion)
	fmt.Println("OS / Arch:", version.OsArch)
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
