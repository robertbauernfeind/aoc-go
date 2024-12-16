package main

import (
	"aoc-in-go/common"
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
	equations := common.SplitToLines(input)

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		equationSum := 0
		for _, equation := range equations {
			splits := strings.Split(equation, ":")
			expectedSum := common.Atoi(splits[0])
			numbers := common.SliceAtoi(strings.Split(strings.Trim(splits[1], " "), " "))

			if isCalculationAMatch(expectedSum, 0, numbers, true) {
				equationSum += expectedSum
			}
		}
		return equationSum
	}
	// solve part 1 here
	equationSum := 0
	for _, equation := range equations {
		splits := strings.Split(equation, ":")
		expectedSum := common.Atoi(splits[0])
		numbers := common.SliceAtoi(strings.Split(strings.Trim(splits[1], " "), " "))

		if isCalculationAMatch(expectedSum, 0, numbers, false) {
			equationSum += expectedSum
		}
	}
	return equationSum
}

// helper function for calculation
func calculate(a, b int, op byte) int {
	calculation := 0
	switch op {
	case '+':
		calculation = a + b
	case '*':
		calculation = a * b
	case '|':
		mul, q := 10, 10
		for q != 0 {
			q = b / mul
			if q > 0 {
				mul *= 10
			}
		}
		calculation = (a * mul) + b
	}

	return calculation
}

func isCalculationAMatch(expectedSum, sum int, input []int, part2 bool) bool {
	if len(input) == 0 {
		return sum == expectedSum
	}

	if sum > expectedSum {
		return false
	}

	if isCalculationAMatch(expectedSum, calculate(sum, input[0], '+'), input[1:], part2) {
		return true
	}

	if part2 && isCalculationAMatch(expectedSum, calculate(sum, input[0], '|'), input[1:], part2) {
		return true
	}

	return isCalculationAMatch(expectedSum, calculate(sum, input[0], '*'), input[1:], part2)
}
