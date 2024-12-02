package day2

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

func TestSafeLevel(t *testing.T) {
	var SafeLevel = []int{7, 6, 4, 2, 1}

	if !LevelIsSafe(SafeLevel) {
		t.Errorf(
			`Test failed, LevelIsSafe(%d) = %t, expected %t`,
			SafeLevel,
			false,
			true,
		)
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
