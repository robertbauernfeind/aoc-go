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
		drXY := drScan2(lines)
		dlXY := dlScan2(lines)
		cnt := 0

		for _, dr := range drXY {
			for _, dl := range dlXY {
				if dl[0] == dr[0] && dl[1] == dr[1] {
					cnt++
				}
			}
		}
		return cnt
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

// part 2
// returns [] x, y
func drScan2(lines []string) [][]int {
	cords := [][]int{}

	// row
	for i := 0; i < len(lines)-2; i++ {
		// col
		for j := 0; j < len(lines[i])-2; j++ {
			mas := string(lines[i][j]) + string(lines[i+1][j+1]) + string(lines[i+2][j+2])

			if isMas(mas) {
				// middle of MAS
				cords = append(cords, []int{j + 1, i + 1})
			}
		}
	}
	return cords
}

// returns [] x, y
func dlScan2(lines []string) [][]int {
	cords := [][]int{}

	// row
	for i := 2; i < len(lines); i++ {
		// col
		for j := 0; j < len(lines[i])-2; j++ {
			mas := string(lines[i][j]) + string(lines[i-1][j+1]) + string(lines[i-2][j+2])

			if isMas(mas) {
				cords = append(cords, []int{j + 1, i - 1})
			}
		}
	}
	return cords
}

func isXmas(s string) bool {
	return s == "XMAS" || s == "SAMX"
}

func isMas(s string) bool {
	return s == "MAS" || s == "SAM"
}
