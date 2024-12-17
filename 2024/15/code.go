package main

import (
	"aoc-in-go/common"
	"fmt"
	"strings"
)

func main() {
	common.RunDev(run)
	// aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	iMap := strings.Split(input, "\n\n")[0]
	iMovements := strings.Split(input, "\n\n")[1]

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	sx, sy := 0, 0
	sIMap := strings.Split(iMap, "\n")
	bMap := [][]byte{}
	for y := range sIMap {
		for x := range sIMap[0] {
			bMap[y][x] = sIMap[y][x]
			if sIMap[y][x] == '@' {
				sx, sy = x, y
			}
		}
	}

	for _, move := range iMovements {
		processMove(sx, sy, byte(move), bMap)
	}

	fmt.Println(sx, sy)
	_ = iMovements
	return 42
}

func processMove(cx, cy int, move byte, bMap [][]byte) (int, int) {
	nx, ny := cx, cy
	switch move {
	case '^':
		ny--
		if collCheck(nx, ny, bMap) {
			panic("Collision")
		}
		// up
	case 'v':
		// down
		ny++
		if collCheck(nx, ny, bMap) {
			panic("Collision")
		}
	case '<':
		// left
		nx--
		if collCheck(nx, ny, bMap) {
			panic("Collision")
		}
	case '>':
		// right
		nx++
		if collCheck(nx, ny, bMap) {
			panic("Collision")
		}
	default:
		panic("Movement not found")
	}

	return nx, ny
}

func collCheck(x, y int, bMap [][]byte) bool {
	return bMap[y][x] == '#'
}
