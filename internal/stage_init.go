package internal

import (
	"fmt"

	testerutils "github.com/codecrafters-io/tester-utils"
)

func stageInit(stageHarness *testerutils.StageHarness) error {
	stageHarness.Logger.Infof("$ echo \"%s\" | ./script.sh -E \"%s\"", "dog", "d")

	result, err := stageHarness.Executable.RunWithStdin([]byte("dog"), "-E", "d")
	if err != nil {
		return err
	}

	if result.ExitCode != 0 {
		return fmt.Errorf("expected exit code %v, got %v", 0, result.ExitCode)
	}

	stageHarness.Logger.Successf("✓ Received exit code %d.", 0)
	return nil
}
