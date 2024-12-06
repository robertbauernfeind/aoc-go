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
	splittedInput := strings.Split(input, "\n")
	labLines := [][]string{}
	for _, v := range splittedInput {
		labLines = append(labLines, strings.Split(v, ""))
	}
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		// simulate initial movement
		sx, sy := getStartingPos(labLines)
		movementMap := predictMovement(sx, sy, labLines)
		simulationMap := [][]string{}
		for _, line := range movementMap {
			simLine := append([]string{}, line...)
			simulationMap = append(simulationMap, simLine)
		}

		possiblePositionCnt := 0

		for fy := 0; fy < len(movementMap); fy++ {
			fmt.Println(movementMap[fy])
			for fx := 0; fx < len(movementMap[0]); fx++ {
				movements := make(map[string]bool)

				if movementMap[fy][fx] == "X" {
					simulationMap[fy][fx] = "#"

					rotation := 0
					posFound := false
					x, y := fx, fy
					for !posFound {
						//x, y, rot
						state := fmt.Sprintf("%d %d %d", fx, fy, rotation)
						if movements[state] {
							possiblePositionCnt++
							posFound = true
						} else {
							movements[state] = true
							// movementMap[y][x] = "X"
							nextX, nextY := predictNextMovement(rotation, x, y)
							if hitsBoundary(simulationMap, nextX, nextY) {
								if simulationMap[nextY][nextX] == "#" {
									rotation = rotate(rotation)
									nextX, nextY = predictNextMovement(rotation, x, y)
								}
								simulationMap[nextY][nextX] = "^"
								x, y = nextX, nextY
							} else {
								fmt.Println("RUN THROUGH")
								// return simulationMap
							}
						}
					}
					simulationMap[y][x] = "."
				}
			}
		}

		return possiblePositionCnt
	}
	// solve part 1 here
	x, y := getStartingPos(labLines)
	movementMap := predictMovement(x, y, labLines)

	return distinctMovements(movementMap)
}

func getStartingPos(labLines [][]string) (x, y int) {
	x, y = -1, -1
	for iy, line := range labLines {
		for ix, c := range line {
			if c == "^" {
				x, y = ix, iy
			}
		}
	}
	if x == -1 || y == -1 {
		panic("Starting pos not found")
	}
	return x, y
}

func predictMovement(sgx, sgy int, labLines [][]string) [][]string {
	movementMap := append([][]string{}, labLines...)
	x, y := sgx, sgy

	// rotation based on clock | 0,3,9,12(0)
	currentRotation := 0

	nextX, nextY := predictNextMovement(currentRotation, x, y)
	for hitsBoundary(labLines, nextX, nextY) {
		movementMap[y][x] = "X"
		nextX, nextY = predictNextMovement(currentRotation, x, y)
		if hitsBoundary(labLines, nextX, nextY) {
			if movementMap[nextY][nextX] == "#" {
				currentRotation = rotate(currentRotation)
				nextX, nextY = predictNextMovement(currentRotation, x, y)
			}
			movementMap[nextY][nextX] = "^"
			x, y = nextX, nextY
		} else {
			return movementMap
		}
	}
	return movementMap
}

func hitsBoundary(labLines [][]string, gx, gy int) bool {
	maxX, maxY := len(labLines[0]), len(labLines)

	if gx >= 0 && gy >= 0 && gx < maxX && gy < maxY {
		return true
	}

	return false
}

func predictNextMovement(currentRotation, gx, gy int) (x, y int) {
	switch currentRotation {
	case 0:
		x, y = gx, gy-1
		break
	case 6:
		x, y = gx, gy+1
		break
	case 3:
		x, y = gx+1, gy
	case 9:
		x, y = gx-1, gy
	default:
		panic("No movement possible")
	}

	return x, y
}

func rotate(rotation int) int {
	rotation += 3
	if rotation > 9 {
		rotation = 0
	}

	return rotation
}

func distinctMovements(labLines [][]string) int {
	cnt := 0
	for _, line := range labLines {
		for _, tile := range line {
			if tile == "X" {
				cnt++
			}
		}
	}

	return cnt
}

func debugPrintMap(lab [][]string) {
	for _, mapLine := range lab {
		fmt.Println(strings.Join(mapLine, ""))
	}
}
