package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type pair[T any] struct {
	first, second T
}

type IntParsingError struct {
	msg string
	err error
}

func PairInts(input []int) (ret []pair[int]) {
	/* Pairs integers two by two in an array and returns the resulting array */

	for idx, elt := range input[:len(input)-1] {
		ret = append(ret, pair[int]{elt, input[idx+1]})
	}

	return ret
}

func Abs(input int) int {
	/* Shortcut to avoid using `math.Abs` */
	if input >= 0 {
		return input
	} else {
		return -1 * input
	}
}

func ReportIsSafe(input []int) bool {
	/* Checks whether a level is safe or not

	Level safety is defined by the maximum difference between two consecutive
	items being below 3 */

	var levelsDiff []int

	for _, pairing := range PairInts(input) {
		levelsDiff = append(levelsDiff, pairing.second-pairing.first)
	}

	// Do a 1-indexed loop to compare values two by two
	for i := 0; i < len(levelsDiff)-1; i++ {
		diff := levelsDiff[i]

		// If both diffs do not share the same sign then level is flucutating, thus unsafe
		if math.Signbit(float64(diff)) != math.Signbit(float64(levelsDiff[i+1])) {
			return false
		}

		if Abs(diff) == 0 || Abs(diff) > 3 {
			return false
		}
	}

	lastReading := Abs(levelsDiff[len(levelsDiff)-1])
	return lastReading != 0 && lastReading <= 3
}

func OpenInputFile(path string) (levels [][]int) {
	/* Opens a file under `path` and reads its lines with the format described in the second day advent of code.

	Returns a list of lists of ints.
	*/

	fp, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		// Input file has one space of separation between ints
		var line = strings.Split(scanner.Text(), " ")

		var inputs []int

		for _, input := range line {
			var parsedInt, err = strconv.ParseInt(input, 10, 64)

			if err != nil {
				panic(IntParsingError{"Int parsing failure", err})
			}

			inputs = append(inputs, int(parsedInt))
		}

		levels = append(levels, inputs)
	}

	return levels
}

func main() {
	levels := OpenInputFile("day2.txt")

	var safeLevels int

	for _, level := range levels {
		safe := ReportIsSafe(level)
		var safeReportsWithTolerance []bool

		for i := 0; i < len(level); i++ {
			copiedLevel := append([]int(nil), level...)

			newSlice := append(copiedLevel[:i], copiedLevel[i+1:]...)

			safeReportsWithTolerance = append(safeReportsWithTolerance, ReportIsSafe(newSlice))
		}

		if safe || slices.Contains(safeReportsWithTolerance, true) {
			safeLevels += 1
		} else {
			fmt.Println("Unsafe level", level)
		}
	}

	fmt.Println("Safe levels count", safeLevels)
}
