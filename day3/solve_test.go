package main

import (
	"testing"
)

// var safelevelstests = map[string]struct {
// 	in  []int
// 	out bool
// }{
// 	"decreasing": {[]int{7, 6, 4, 2, 1}, true},
// 	"increasing": {[]int{1, 3, 6, 7, 9}, true},
// }

// func TestSafeLevel(t *testing.T) {
// 	for name, testcase := range safelevelstests {
// 		t.Run(string(name), func(t *testing.T) {
// 			if ReportIsSafe(testcase.in) != testcase.out {
// 				t.Errorf(
// 					`Test failed, LevelIsSafe(%d) = %t, expected %t`,
// 					testcase.in,
// 					false,
// 					true,
// 				)
// 			}
// 		})
// 	}
// }

var basicMulInputs = map[string]struct {
	in  string
	out int
}{
	"single_digits": {"mul(2,3)", 2 * 3},
	"double_digits": {"mul(20,30)", 20 * 30},
	"mixed_digits":  {"mul(1337,42)", 1337 * 42},
}

func TestBasic(t *testing.T) {
	for name, testcase := range basicMulInputs {
		t.Run(string(name), func(t *testing.T) {
			result := ParseInput(testcase.in)

			if result != testcase.out {
				t.Errorf(
					`Test failed, ParseInput(%s) = %d, expected %d`,
					testcase.in,
					result,
					testcase.out,
				)
			}
		})
	}
}

var invalidInputs = map[string]struct {
	in  string
	out int
}{
	"missing_parenthesis": {"mul(2,3", 0},
	"missing_operator":    {"(20,30)", 0},
	"spaced_operands":     {"mul( 1337 , 42 )", 0},
}

func TestInvalidInputs(t *testing.T) {
	for name, testcase := range invalidInputs {
		t.Run(string(name), func(t *testing.T) {
			result := ParseInput(testcase.in)

			if result != testcase.out {
				t.Errorf(
					`Test failed, ParseInput(%s) = %d, expected %d`,
					testcase.in,
					result,
					testcase.out,
				)
			}
		})
	}
}

func TestMixedInput(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := 2*4 + 5*5 + 11*8 + 8*5

	result := ParseInput(input)

	if result != expected {
		t.Errorf(
			`Test failed, ParseInput(%s) = %d, expected %d`,
			input,
			result,
			expected,
		)
	}

}
