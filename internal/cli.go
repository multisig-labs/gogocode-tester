package internal

import (
	"fmt"
	"time"

	tester_utils "github.com/codecrafters-io/tester-utils"
	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

func RunTestCases(exePath string, stage string) error {
	exec := executable.NewExecutable(exePath)
	isDebug := false
	for _, testCase := range testCases {

		stageHarness := tester_utils.StageHarness{
			Logger:     logger.GetLogger(isDebug, fmt.Sprintf("[%s] ", "prefix")),
			Executable: exec,
		}

		logger := stageHarness.Logger
		logger.Infof("Running tests for %s", testCase.Slug)

		stepResultChannel := make(chan error, 1)
		go func() {
			err := testCase.TestFunc(&stageHarness)
			stepResultChannel <- err
		}()

		timeout := testCase.CustomOrDefaultTimeout()

		var err error
		select {
		case stageErr := <-stepResultChannel:
			err = stageErr
		case <-time.After(timeout):
			err = fmt.Errorf("timed out, test exceeded %d seconds", int64(timeout.Seconds()))
		}

		if err != nil {
			logger.Errorf("%s", err)
		} else {
			logger.Successf("Test passed.")
		}

		stageHarness.RunTeardownFuncs()

		if err != nil {
			return err
		}
	}

	return nil
}
