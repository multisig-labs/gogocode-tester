package testtools

type TesterDefinition struct {
	ExecutableFileName string
	TestCases          []TestCase
}

func (t TesterDefinition) TestCaseBySlug(slug string) TestCase {
	for _, testCase := range t.TestCases {
		if testCase.Slug == slug {
			return testCase
		}
	}

	return TestCase{}
}
