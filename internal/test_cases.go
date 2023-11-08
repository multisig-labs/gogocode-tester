package internal

import (
	testerutils "github.com/codecrafters-io/tester-utils"
)

var testCases = []testerutils.TestCase{
	{
		Slug:     "init",
		TestFunc: stageInit,
	},
	{
		Slug:     "init2",
		TestFunc: stageInit,
	},
}
