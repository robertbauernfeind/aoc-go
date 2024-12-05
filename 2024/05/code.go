package main

import (
	"math"
	"slices"
	"strconv"
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
	printSplits := strings.Split(input, "\n\n")
	printOrders := strings.Split(printSplits[0], "\n")
	printUpdates := strings.Split(printSplits[1], "\n")

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	middleSum := 0
	for _, v := range printUpdates {
		updatedPages := strings.Split(v, ",")
		if correctUpdate(updatedPages, printOrders) {
			middleIdx := int(math.Round(float64(len(updatedPages))/2) - 1)
			val, _ := strconv.Atoi(updatedPages[middleIdx])
			middleSum += val
		}
	}
	return middleSum
}

func correctUpdate(u []string, printOrders []string) bool {
	for i, p1 := range u {
		pagesBefore := u[:i]
		pagesAfer := u[i+1:]

		for _, b := range pagesBefore {
			bef := b + "|" + p1
			idx := slices.Index(printOrders, bef)
			if idx == -1 {
				return false
			}
		}

		for _, a := range pagesAfer {
			aft := p1 + "|" + a
			idx := slices.Index(printOrders, aft)
			if idx == -1 {
				return false
			}
		}
		_, _, _ = p1, pagesAfer, pagesBefore
	}
	return true
}
