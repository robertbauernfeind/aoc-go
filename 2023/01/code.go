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
	lines := strings.Split(input, "\n")
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	sum := 0
	for _, line := range lines {
		nums := [2]string{}
		for _, c := range line {
			if _, err := strconv.Atoi(string(c)); err != nil {
				continue
			}
			nums[0] = string(c)
		}

		for i := 0; i < len(line); i++ {
			c := string(line[len(line)-i-1])
			if _, err := strconv.Atoi(string(c)); err != nil {
				continue
			}
			nums[1] = string(c)
		}
		numStr := nums[1] + nums[0]
		num, _ := strconv.Atoi(numStr)
		sum += num
	}
	return sum
}