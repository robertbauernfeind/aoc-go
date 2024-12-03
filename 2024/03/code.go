package main

import (
	"regexp"
	"strconv"

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
	sumProd := 0
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		re := regexp.MustCompile(`(do\(\)|don't\(\))|mul\((\d{1,3}),(\d{1,3})\)`)
		matches := re.FindAllStringSubmatch(input, -1)

		// initialy set true because first occurance of `mul(x,y)` is always enabled without a `do()`
		enabled := true
		for _, v := range matches {
			if len(v[1]) > 0 {
				if v[1] == "do()" {
					enabled = true
				} else if v[1] == "don't()" {
					enabled = false
				}
			} else if len(v[2]) > 0 && len(v[3]) > 0 && enabled {
				sumProd += atoi(v[2]) * atoi(v[3])
			}
		}
		return sumProd
	}
	// solve part 1 here
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)

	for _, v := range matches {
		sumProd += atoi(v[1]) * atoi(v[2])
	}
	return sumProd
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
