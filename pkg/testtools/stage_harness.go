package testtools

// From github.com/codecrafters-io

import (
	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/multisig-labs/gogocode-tester-sample/pkg/logger"
)

type StageHarness struct {
	// Logger is to be used for all logs generated from the test function.
	Logger *logger.Logger

	// Executable is the program to be tested.
	Executable *executable.Executable

	// teardownFuncs are run once the error has been reported to the user
	teardownFuncs []func()
}

func (s *StageHarness) RegisterTeardownFunc(teardownFunc func()) {
	s.teardownFuncs = append(s.teardownFuncs, teardownFunc)
}

func (s StageHarness) RunTeardownFuncs() {
	for _, teardownFunc := range s.teardownFuncs {
		teardownFunc()
	}
}
