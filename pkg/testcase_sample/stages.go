package testcase_sample

import (
	"bytes"
	"fmt"

	"github.com/multisig-labs/gogocode-tester-sample/pkg/testtools"
)

func stageEchoString(stageHarness *testtools.StageHarness) error {
	stageHarness.Logger.Infof("$ ./run.sh \"%s\"", "dog")

	result, err := stageHarness.Executable.Run("dog")
	if err != nil {
		return err
	}

	if result.ExitCode != 0 {
		return fmt.Errorf("expected exit code %v, got %v", 0, result.ExitCode)
	}

	if !bytes.Equal(result.Stdout, []byte("dog")) {
		return fmt.Errorf("expected 'dog', got '%s'", string(result.Stdout))
	}

	stageHarness.Logger.Successf("✓ Received exit code %d.", 0)
	return nil
}

func stageEchoModifiedString(stageHarness *testtools.StageHarness) error {
	stageHarness.Logger.Infof("$ ./run.sh \"%s\"", "cat")

	result, err := stageHarness.Executable.Run("cat")
	if err != nil {
		return err
	}

	if result.ExitCode != 0 {
		return fmt.Errorf("expected exit code %v, got %v", 0, result.ExitCode)
	}

	if !bytes.Equal(result.Stdout, []byte("dog")) {
		return fmt.Errorf("expected 'dog', got '%s'", string(result.Stdout))
	}

	stageHarness.Logger.Successf("✓ Received exit code %d.", 0)
	return nil
}
