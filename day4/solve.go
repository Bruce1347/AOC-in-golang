package day4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseInputLinesHorizontally(input []string) (res int) {
	for _, line := range input {
		if strings.Contains(line, "XMAS") || strings.Contains(line, "SAMX") {
			res += 1
		}
	}
	return res
}

func ParseInputLinesVertically(input []string) (res int) {
	nbcols := len(input[0])
	nbrows := len(input)

	for j := 0; j < nbcols; j++ {
		var buf strings.Builder

		for i := 0; i < nbrows; i++ {
			buf.WriteString(string(input[i][j]))
		}

		bufstr := buf.String()

		if strings.Contains(bufstr, "XMAS") || strings.Contains(bufstr, "SAMX") {
			res += 1
		}
	}

	return res
}

func ParseInputLinesDiagonal(input []string) (res int) {
	nbcols := len(input[0])
	nbrows := len(input)

	for i := 0; i < nbrows; i++ {
		// Avoid out of bounds with a lookahead of 3
		for j := 0; j+3 < nbcols && i+3 < nbrows; j++ {
			var buf strings.Builder
			buf.WriteString(string(input[i][j]))
			buf.WriteString(string(input[i+1][j+1]))
			buf.WriteString(string(input[i+2][j+2]))
			buf.WriteString(string(input[i+3][j+3]))

			bufstr := buf.String()

			if strings.Contains(bufstr, "XMAS") || strings.Contains(bufstr, "SAMX") {
				res += 1
			}
		}
	}

	for i := nbrows - 1; i >= 0; i-- {
		for j := 0; j+3 < nbcols && i-3 >= 0; j++ {
			var bufReverse strings.Builder
			bufReverse.WriteString(string(input[i][j]))
			bufReverse.WriteString(string(input[i-1][j+1]))
			bufReverse.WriteString(string(input[i-2][j+2]))
			bufReverse.WriteString(string(input[i-3][j+3]))

			bufreverse := bufReverse.String()

			if strings.Contains(bufreverse, "XMAS") || strings.Contains(bufreverse, "SAMX") {
				res += 1
			}
		}
	}

	return res
}

func ParseInput(input []string) int {
	return ParseInputLinesHorizontally(input) + ParseInputLinesVertically(input) + ParseInputLinesDiagonal(input)
}

func main() {
	input, err := os.ReadFile("day4.txt")

	lines := strings.Split(string(input), "\n")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result", ParseInput(lines))
}
