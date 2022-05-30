package demo

import (
	"testing"
)

func TestPlus(t *testing.T) {
	// Arrange - prepare test cases, input arguments and expected result.
	testCase := struct{ x, y, expected int }{2, 1, 3}

	// Act - call function and pass test case arguments
	result := Plus(testCase.x, testCase.y) // test function Plus() in demo.go

	// Assert - compare actual result with expected result
	if result != testCase.expected {
		t.Errorf("Plus(%d, %d) == %d, but expect %d",
			testCase.x, testCase.y, result, testCase.expected)
	}
}

func TestMinus(t *testing.T) {
	testCases := []struct {
		x, y, expected int
	}{
		{2, 1, 1},    // test case 1
		{3, 1, 2},    // test case 2
		{100, 1, 99}, // test case 3
	}

	for _, testCase := range testCases { // execute test cases one by one
		result := Minus(testCase.x, testCase.y)
		if result != testCase.expected {
			t.Errorf("Minus(%d, %d) == %d, but expect %d",
				testCase.x, testCase.y, result, testCase.expected)
		}
	}

}

func TestAddBracket(t *testing.T) {
	testCases := []struct {
		s, expected string
	}{
		{"hello world", "[hello world]"},
		{"Nice", "[Nice]"},
	}

	for _, testCase := range testCases {
		result := AddBracket(testCase.s)
		if result != testCase.expected {
			t.Errorf("AddBracket(%s) == %s, but expect %s ",
				testCase.s, result, testCase.expected)
		}
	}
}
