package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ParseInput(input string) (res int) {
	mul_regexp := regexp.MustCompile("(mul\\([0-9]+,[0-9]+\\))")
	// Same regexp as before but with capturing groups for operands
	ints_regexp := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)")

	for _, match := range mul_regexp.FindAllString(input, -1) {
		operands := ints_regexp.FindAllStringSubmatch(match, -1)

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
	}

	return res
}

func main() {

	file, err := os.Open("day3.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	acc := 0

	for scanner.Scan() {
		acc += ParseInput(scanner.Text())
	}

	fmt.Println("Result", acc)

}
