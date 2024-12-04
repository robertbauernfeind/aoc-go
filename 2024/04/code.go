// 04/code.go
package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	// common.RunDev(run)
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	lines := strings.Split(input, "\n")
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	cnt := hScan(lines)
	cnt += vScan(lines)
	cnt += drScan(lines)
	cnt += dlScan(lines)
	return cnt
}

func hScan(lines []string) int {
	cnt := 0
	for _, line := range lines {
		for i := 3; i < len(line); i++ {
			xmas := string(line[i-3]) + string(line[i-2]) + string(line[i-1]) + string(line[i])
			if isXmas(xmas) {
				cnt++
			}
		}
	}
	return cnt
}

func vScan(lines []string) int {
	cnt := 0
	for i := 3; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			xmas := string(lines[i-3][j]) + string(lines[i-2][j]) + string(lines[i-1][j]) + string(lines[i][j])
			if isXmas(xmas) {
				cnt++
			}
		}
	}

	return cnt
}

func drScan(lines []string) int {
	cnt := 0

	// row
	for i := 0; i < len(lines)-3; i++ {
		// col
		for j := 0; j < len(lines[i])-3; j++ {
			xmas := string(lines[i][j]) + string(lines[i+1][j+1]) + string(lines[i+2][j+2]) + string(lines[i+3][j+3])

			if isXmas(xmas) {
				cnt++
			}
		}
	}
	return cnt
}

func dlScan(lines []string) int {
	cnt := 0

	// row
	for i := 3; i < len(lines); i++ {
		// col
		for j := 0; j < len(lines[i])-3; j++ {
			xmas := string(lines[i][j]) + string(lines[i-1][j+1]) + string(lines[i-2][j+2]) + string(lines[i-3][j+3])

			if isXmas(xmas) {
				cnt++
			}
		}
	}
	return cnt
}

func isXmas(s string) bool {
	return s == "XMAS" || s == "SAMX"
}
