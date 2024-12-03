package main

import (
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
	reports := strings.Split(input, "\n")
	safeCtr := 0
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		for _, report := range reports {
			levels := stringSliceToIntSlice(strings.Split(report, " "))
			if safe(levels) || safeWithRemoving(levels) {
				safeCtr++
			}
		}
		return safeCtr
	}
	// solve part 1 here
	for _, report := range reports {
		levels := stringSliceToIntSlice(strings.Split(report, " "))
		if safe(levels) {
			safeCtr++
		}
	}
	return safeCtr
}

func safeWithRemoving(levels []int) bool {
	for i := range levels {
		nl := append([]int{}, levels[:i]...)
		nl = append(nl, levels[i+1:]...)

		if safe(nl) {
			return true
		}
	}
	return false
}

func safe(levels []int) bool {
	// safe when either asscendin or descending
	// and diff between levels is 1 or 2
	asc := levels[0] < levels[1]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		// if diff is less then 0, levels are descending
		// if diff is equal to 0 its neither asc or desc so its false
		if diff > 0 != asc {
			return false
		}

		if diff < 0 {
			diff *= -1
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func stringSliceToIntSlice(s []string) []int {
	var r []int
	for _, i := range s {
		j, _ := strconv.Atoi(i)
		r = append(r, j)
	}
	return r
}
