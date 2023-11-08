package testcase_sample

import "github.com/multisig-labs/gogocode-tester-sample/pkg/testtools"

// //go:generate stringer -type=Stage
// type Stage uint8

// const (
// 	Init Stage = iota
// 	NextOne
//  Etc
// )

const CourseSlug = "sample"

var testCases = []testtools.TestCase{
	{
		Slug:     "echo-string",
		TestFunc: stageEchoString,
	},
	{
		Slug:     "echo-modified-string",
		TestFunc: stageEchoModifiedString,
	},
}

// func GetNextStage(slug string) (string, error) {
// 	for i, testCase := range testCases {
// 		if testCase.Slug == slug {
// 			// Check if the current test case is not the last one
// 			if i < len(testCases)-1 {
// 				return testCases[i+1].Slug, nil
// 			}
// 			// We are at the end, return empty string
// 			return "", nil
// 		}
// 	}
// 	return "", fmt.Errorf("next slug not found")
// }

// Return all test cases up to the one specified (empty string for all)
// func casesToRun(slug string) []*testtools.TestCase {
// 	out := []*testtools.TestCase{}
// 	for _, tc := range testCases {
// 		fu := tc
// 		out = append(out, &fu)
// 		if slug == tc.Slug {
// 			break
// 		}
// 	}
// 	return out
// }
