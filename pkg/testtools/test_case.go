package testtools

import "time"

// From github.com/codecrafters-io

// TestCase represents a test case that'll be run against the user's code.
type TestCase struct {
	// Slug is the unique identifier for this test case.
	Slug string

	// TestFunc is the function that'll be run against the user's code.
	TestFunc func(stageHarness *StageHarness) error

	// Timeout is the maximum amount of time that the test case can run for.
	Timeout time.Duration
}

func (t TestCase) CustomOrDefaultTimeout() time.Duration {
	if (t.Timeout == 0) || (t.Timeout == time.Duration(0)) {
		return 10 * time.Second
	} else {
		return t.Timeout
	}
}
