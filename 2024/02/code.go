package main

import (
	"math"
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
	safe := 0
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return 12
	}
	// solve part 1 here
	for _, report := range reports {
		levels := stringSliceToIntSlice(strings.Split(report, " "))
		if (isIncreasing(levels) || isDecreasing(levels)) && diffIsSave(levels) {
			safe++
		}
	}
	return safe
}

func stringSliceToIntSlice(s []string) []int {
	var r []int
	for _, i := range s {
		j, _ := strconv.Atoi(i)
		r = append(r, j)
	}
	return r
}

func isIncreasing(levels []int) bool {
	for i := 0; i < len(levels)-1; i++ {
		lvl1, lvl11 := levels[i], levels[i+1]
		if lvl1 >= lvl11 {
			return false
		}
	}

	return true
}

func isDecreasing(levels []int) bool {
	for i := 0; i < len(levels)-1; i++ {
		lvl1, lvl11 := levels[i], levels[i+1]
		if lvl1 <= lvl11 {
			return false
		}
	}

	return true
}

func diffIsSave(levels []int) bool {
	for i := 0; i < len(levels)-1; i++ {
		lvl1, lvl11 := levels[i], levels[i+1]
		diff := absDiff(lvl1, lvl11)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func absDiff(x, y int) int {
	return int(math.Abs(float64(x - y)))
}
