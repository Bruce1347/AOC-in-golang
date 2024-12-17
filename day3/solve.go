package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ParseInput(input string) (res int) {
	mul_regexp := regexp.MustCompile(`mul\(\d+,\d+\)|do[n't]*\(\)`)
	// Same regexp as before but with capturing groups for operands
	ints_regexp := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	count_enabled := true

	for _, match := range mul_regexp.FindAllString(input, -1) {
		operands := ints_regexp.FindAllStringSubmatch(match, -1)

		if match == "do()" || match == "don't()" {
			switch match {
			case "do()":
				count_enabled = true
			case "don't()":
				count_enabled = false
			}
			continue
		}

		if count_enabled {
			if len(operands) == 0 {
				// Something wrong happened
				return -1
			}

			a, err_a := strconv.Atoi(operands[0][1])
			b, err_b := strconv.Atoi(operands[0][2])

			if err_a != nil || err_b != nil {
				panic("Int conversion failure")
			}

			res += a * b

			fmt.Println(res)
		}

	}

	return res
}

func main() {

	input, err := os.ReadFile("day3.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result", ParseInput(string(input)))

}
