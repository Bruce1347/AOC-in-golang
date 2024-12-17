package day4

import "testing"

func TestParseSimpleLine(t *testing.T) {
	input := []string{"MMMSXXMASM"}
	res := ParseInput(input)

	if res != 1 {
		t.Errorf(
			`Test failed, ParseInput(%s) = %d, expected %d`,
			input,
			res,
			1,
		)
	}
}

func TestParseVerticalMatch(t *testing.T) {
	input := []string{
		".X....",
		".M....",
		".A....",
		".S....",
		"......",
	}
	res := ParseInput(input)

	if res != 1 {
		t.Errorf(
			`Test failed, ParseInput(%s) = %d, expected %d`,
			input,
			res,
			1,
		)
	}
}

func TestParseDiagonalMatch(t *testing.T) {
	input := []string{
		"X.....",
		".M....",
		"..A...",
		"...S..",
		"......",
	}
	res := ParseInput(input)

	if res != 1 {
		t.Errorf(
			`Test failed, ParseInput(%s) = %d, expected %d`,
			input,
			res,
			1,
		)
	}
}

var diagonalParseInputs = map[string]struct {
	in  []string
	out int
}{
	"diagonal": {
		[]string{
			"X.....",
			".M....",
			"..A...",
			"...S..",
			"......",
		},
		1,
	},
	"diagonal_offset": {
		[]string{
			"..X...",
			"...M..",
			"....A.",
			".....S",
			"......",
		},
		1,
	},
	"diagonal_offset_reversed": {
		[]string{
			"..S...",
			"...A..",
			"....M.",
			".....X",
			"......",
		},
		1,
	},
	"diagonal_reversed": {
		[]string{
			"S.....",
			".A....",
			"..M...",
			"...X..",
			"......",
		},
		1,
	},
	"diagonal_bottom": {
		[]string{
			"......",
			"...S..",
			"..A...",
			".M....",
			"X.....",
		},
		1,
	},
	"diagonal_bottom_reversed": {
		[]string{
			"......",
			"...X..",
			"..M...",
			".A....",
			"S.....",
		},
		1,
	},
	"diagonal_bottom_right": {
		[]string{
			"......",
			"..X...",
			"...M..",
			"....A.",
			".....S",
		},
		1,
	},
	"diagonal_bottom_right_reversed": {
		[]string{
			"......",
			"..S...",
			"...A..",
			"....M.",
			".....X",
		},
		1,
	},
	"diagonal_bottom_right_reversed_offset": {
		[]string{
			"......",
			"S.....",
			".A....",
			"..M...",
			"...X..",
		},
		1,
	},
}

func TestDiagonal(t *testing.T) {
	for name, testcase := range diagonalParseInputs {
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

func TestParseComplexCase(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	res := ParseInput(input)

	if res != 18 {
		t.Errorf(
			`Test failed, ParseInput(%s) = %d, expected %d`,
			input,
			res,
			18,
		)
	}

}
