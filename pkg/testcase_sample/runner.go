package testcase_sample

import (
	"fmt"
	"time"

	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/multisig-labs/gogocode-tester-sample/pkg/logger"
	"github.com/multisig-labs/gogocode-tester-sample/pkg/testtools"
)

// Run all stages until a failure occurs.
// Expects to
// Return the slug of the last test to pass
func RunTestCases(exePath string) (string, error) {
	exec := executable.NewExecutable(exePath)
	// TODO if we need this for logging then need way for user to set debug = true
	isDebug := false
	logger.GetLogger(false, "").Plainln(testtools.AsciiArt)
	lastSuccessfulSlug := ""
	for _, testCase := range testCases {
		log := logger.GetLogger(isDebug, fmt.Sprintf("[%s] ", testCase.Slug))
		stageHarness := testtools.StageHarness{
			Logger:     log,
			Executable: exec,
		}

		log.Infof("Running tests for course %s, stage %s", CourseSlug, testCase.Slug)

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
			log.Errorf("%s", err)
		} else {
			log.Successf("Test passed.")
			lastSuccessfulSlug = testCase.Slug
		}

		stageHarness.RunTeardownFuncs()

		if err != nil {
			return lastSuccessfulSlug, err
		}
	}

	return lastSuccessfulSlug, nil
}
