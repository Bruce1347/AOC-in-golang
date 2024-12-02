package main

import (
	"reflect"
	"testing"
)

func TestPairIntsSimple(t *testing.T) {
	var input = []int{1, 2, 3, 4}
	var expected = []pair{pair{1, 2}, pair{2, 3}, pair{3, 4}}

	if !reflect.DeepEqual(PairInts(input), expected) {
		t.Errorf(
			`Test failed, PairInts(%d) = %d, expected %d`,
			input,
			PairInts(input),
			expected,
		)
	}
}

func TestPairIntsOddCount(t *testing.T) {
	var input = []int{1, 2, 3, 4, 5}
	var expected = []pair{pair{1, 2}, pair{2, 3}, pair{3, 4}, pair{4, 5}}

	if !reflect.DeepEqual(PairInts(input), expected) {
		t.Errorf(
			`Test failed, PairInts(%d) = %d, expected %d`,
			input,
			PairInts(input),
			expected,
		)
	}
}

var safelevelstests = map[string]struct {
	in  []int
	out bool
}{
	"decreasing": {[]int{7, 6, 4, 2, 1}, true},
	"increasing": {[]int{1, 3, 6, 7, 9}, true},
}

func TestSafeLevel(t *testing.T) {
	for name, testcase := range safelevelstests {
		t.Run(string(name), func(t *testing.T) {
			if LevelIsSafe(testcase.in) != testcase.out {
				t.Errorf(
					`Test failed, LevelIsSafe(%d) = %t, expected %t`,
					testcase.in,
					false,
					true,
				)
			}
		})
	}
}

func TestUnSafeLevel(t *testing.T) {
	var SafeLevel = []int{69, 7, 6, 4, 2, 1}

	if LevelIsSafe(SafeLevel) {
		t.Errorf(
			`Test failed, LevelIsSafe(%d) = %t, expected %t`,
			SafeLevel,
			false,
			true,
		)
	}
}

func TestUnsafeLevelDecrease(t *testing.T) {
	var UnsafeLevel = []int{1, 2, 3, 2, 4}

	if LevelIsSafe(UnsafeLevel) {
		t.Errorf(
			`Test failed, LevelIsSafe(%d) = %t, expected %t`,
			UnsafeLevel,
			false,
			true,
		)

	}
}

func TestUnsafeLevelDecreaseTooHigh(t *testing.T) {
	var UnsafeLevel = []int{9, 7, 6, 2, 1}

	if LevelIsSafe(UnsafeLevel) {
		t.Errorf(
			`Test failed, LevelIsSafe(%d) = %t, expected %t`,
			UnsafeLevel,
			false,
			true,
		)

	}
}

func TestUnsafeLevelIncreaseTooHigh(t *testing.T) {
	var UnsafeLevel = []int{83, 87, 89, 90, 93, 96, 99}

	if LevelIsSafe(UnsafeLevel) {
		t.Errorf(
			`Test failed, LevelIsSafe(%d) = %t, expected %t`,
			UnsafeLevel,
			false,
			true,
		)

	}
}
