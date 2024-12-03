package main

import (
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
		occ := map[int]int{}

		for _, v := range right {
			occ[v]++
		}

		for _, v := range left {
			distance += occ[v] * v
		}
		return distance
	}
	// solve part 1 here

	slices.Sort(left)
	slices.Sort(right)

	for i, v := range left {
		localDist := v - right[i]
		distance += abs(localDist)
	}

	return distance
}

func splitToSlices(s string) (l []int, r []int) {
	inputSlice := strings.Split(s, "\n")
	for _, v := range inputSlice {
		splits := strings.Split(v, " ")

		n1, _ := strconv.Atoi(splits[0])
		n2, _ := strconv.Atoi(splits[3])
		l = append(l, n1)
		r = append(r, n2)
	}

	return l, r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
