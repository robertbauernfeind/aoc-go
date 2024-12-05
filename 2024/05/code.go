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

	corr, incorr := [][]string{}, [][]string{}

	for _, v := range printUpdates {
		updatedPages := strings.Split(v, ",")
		if updateIsCorrect(updatedPages, printOrders) {
			corr = append(corr, updatedPages)
		} else {
			incorr = append(incorr, updatedPages)
		}
	}

	middleSum := 0
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		for _, v := range incorr {
			if correctUpdate(v, printOrders) {
				middleIdx := int(math.Round(float64(len(v))/2) - 1)
				val, _ := strconv.Atoi(v[middleIdx])
				middleSum += val
			}
		}
		return middleSum
	}
	// solve part 1 here
	for _, v := range corr {

		middleIdx := int(math.Round(float64(len(v))/2) - 1)
		val, _ := strconv.Atoi(v[middleIdx])
		middleSum += val
	}
	return middleSum
}

func updateIsCorrect(u []string, printOrders []string) bool {
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
	}
	return true
}

func correctUpdate(u []string, printOrders []string) bool {
	timeout := 0
	// timeout to stop for loop if something goes wrong
	for !updateIsCorrect(u, printOrders) && timeout <= 20 {
		for i, p1 := range u {
			pagesBefore := u[:i]
			pagesAfer := u[i+1:]

			for _, b := range pagesBefore {
				bef := b + "|" + p1
				idx := slices.Index(printOrders, bef)
				if idx == -1 {
					bef := p1 + "|" + b
					idx = slices.Index(printOrders, bef)

					if idx == -1 {
						return false
					} else {
						bIdx := slices.Index(u, b)
						u[i], u[bIdx] = u[bIdx], u[i]
					}
				}
			}

			for _, a := range pagesAfer {
				aft := p1 + "|" + a
				idx := slices.Index(printOrders, aft)
				if idx == -1 {
					bef := a + "|" + p1
					idx = slices.Index(printOrders, bef)

					if idx == -1 {
						return false
					} else {
						aIdx := slices.Index(u, a)
						u[i], u[aIdx] = u[aIdx], u[i]
					}
				}
			}
		}

		timeout++
	}
	return true
}
