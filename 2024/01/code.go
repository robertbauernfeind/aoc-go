package main

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	left, right := splitToSlices(input)
	var distance int = 0
	if part2 {
		// solve part 2 here

		for _, v := range left {
			occurences := occurences(right, v)
			distance += occurences * v
		}
		return distance
	}
	// solve part 1 here

	slices.Sort(left)
	slices.Sort(right)

	for i := 0; i < len(left); i++ {
		var localDist float64 = float64(left[i]) - float64(right[i])
		distance += int(math.Abs(localDist))
	}

	return distance
}

func convToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func splitToSlices(s string) (left []int, right []int) {
	inputSlice := strings.Split(s, "\n")
	_, _ = left, right
	for _, v := range inputSlice {
		splits := strings.Split(v, " ")

		left = append(left, convToInt(strings.Trim(splits[0], " ")))
		right = append(right, convToInt(strings.Trim(splits[3], " ")))
	}

	return left, right
}

func occurences(slice []int, value int) int {
	occurences := 0
	for _, v := range slice {
		if v == value {
			occurences++
		}
	}

	return occurences
}
