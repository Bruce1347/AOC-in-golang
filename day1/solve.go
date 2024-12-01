package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func MergeArrays(arr1 []int, arr2 []int) []int {
	// No merges need
	if len(arr1) == 0 {
		return arr2
	}

	if len(arr2) == 0 {
		return arr1
	}

	// Compare heads and reiterate with the lowest int as the first element
	if arr1[0] < arr2[0] {
		return append(
			[]int{arr1[0]},
			MergeArrays(arr1[1:], arr2)...,
		)
	} else {
		return append(
			[]int{arr2[0]},
			MergeArrays(arr1, arr2[1:])...,
		)
	}
}

func MergeSort(arr []int) []int {
	// End iteration
	if len(arr) <= 1 {
		return arr
	}

	// Find the pivot index
	midIdx := int(len(arr) / 2)

	// Merge the recursive calls
	return MergeArrays(
		// Split and sort array from 0 (inclusive) to midIdx (exclusive)
		MergeSort(arr[:midIdx]),
		// Split and sort array from midIdx (inclusive) to the end
		MergeSort(arr[midIdx:]),
	)
}

// Emulate 2-lengthed tuples because i'm a python nerd
type Pair[T any] struct {
	First  T
	Second T
}

func OpenInputFile(path string) Pair[[]int] {
	/* Opens a file under `path` and reads its contents to populate two separated
	arrays as described in the first advent of code exercise.

	Returns a `Pair` of these arrays
	*/
	fp, err := os.Open(path)

	if err != nil {
		// Something wrong happened, exit.
		panic(err)
	}

	// Do not forget to close the file pointer at the end.
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var arr1 []int
	var arr2 []int

	for scanner.Scan() {
		// Input file has 3 spaces of separation
		var inputs = strings.Split(scanner.Text(), "   ")

		var firstInt, err1 = strconv.ParseInt(inputs[0], 10, 64)
		var secondInt, err2 = strconv.ParseInt(inputs[1], 10, 64)

		if err1 != nil {
			panic(err1)
		}

		if err2 != nil {
			panic(err2)
		}

		arr1 = append(arr1, int(firstInt))
		arr2 = append(arr2, int(secondInt))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return Pair[[]int]{arr1, arr2}
}

func ComputeDistance(firstArray []int, secondArray []int) int {
	if len(firstArray) != len(secondArray) {
		panic("Arrays length mismatch")
	}

	var acc int = 0

	for i := 0; i < len(firstArray); i++ {
		var first = firstArray[i]
		var second = secondArray[i]

		// second may be inferior to first but their distance won't change,
		// using the absolute value avoids subtracting values
		acc = acc + int(math.Abs(float64(first-second)))
	}
	return acc
}

func ComputeSimilarity(firstArray []int, secondArray []int) int {
	var ret int = 0

	for firstArrayIdx := 0; firstArrayIdx < len(firstArray); firstArrayIdx++ {
		var elt int = firstArray[firstArrayIdx]

		var eltCount int = 0

		for secondArrayIdx := 0; secondArrayIdx < len(secondArray); secondArrayIdx++ {
			if secondArray[secondArrayIdx] == elt {
				eltCount = eltCount + 1
			}
		}

		ret = ret + elt*eltCount
	}

	return ret
}

func main() {
	var arrs = OpenInputFile("day1.txt")

	// Sort the arrays
	var firstArray = MergeSort(arrs.First)
	var secondArray = MergeSort(arrs.Second)

	var acc = ComputeDistance(firstArray, secondArray)

	fmt.Println(acc)
	fmt.Println(ComputeSimilarity(firstArray, secondArray))

}
