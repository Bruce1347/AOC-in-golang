package day2

type pair struct {
	first, second int
}

func PairInts(input []int) (ret []pair) {
	/* Pairs integers two by two in an array and returns the resulting array */

	for idx, elt := range input[:len(input)-1] {
		ret = append(ret, pair{elt, input[idx+1]})
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

func LevelIsSafe(input []int) bool {
	/* Checks whether a level is safe or not

	Level safety is defined by the maximum difference between two consecutive
	items being below 3 */

	var levelsDiff []int

	for _, pairing := range PairInts(input) {
		levelsDiff = append(levelsDiff, Abs(pairing.first-pairing.second))
	}

	for _, diff := range levelsDiff {
		if diff > 3 {
			return false
		}
	}

	return true
}

func main() {

}
