package main

import (
	"aoc-in-go/common"
	"fmt"
	"strings"
)

func main() {
	common.RunDev(run)
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}
	formattedDisk := formatDisk(input)
	ret := move(formattedDisk)
	return checkSum(ret)
}

func formatDisk(disk string) string {
	var formattedDisk strings.Builder
	blockId := 0
	for i := 0; i < len(disk); i++ {
		conv := int(disk[i] - '0') // Convert char to int
		if i%2 == 0 {              // If even index, it's a file block
			for j := 0; j < conv; j++ {
				formattedDisk.WriteString(fmt.Sprintf("%d", blockId))
			}
			blockId++
		} else { // If odd index, it's free space
			for j := 0; j < conv; j++ {
				formattedDisk.WriteString(".")
			}
		}
	}
	return formattedDisk.String()
}

func move(fdisk string) string {
	chars := strings.Split(fdisk, "")
	for i := 0; i < len(fdisk); i++ {
		revIdx := len(fdisk) - 1 - i
		freeSpace := findFreeSpace(chars)
		if freeSpace > revIdx {
			break
		}
		// Swap the free space with the last file block
		chars[freeSpace], chars[revIdx] = chars[revIdx], chars[freeSpace]
	}
	return strings.Join(chars, "")
}

func findFreeSpace(chars []string) int {
	for i, ch := range chars {
		if ch == "." {
			return i
		}
	}
	return -1 // This should not happen if the input is well-formed
}

func checkSum(fdisk string) int {
	checkSum := 0
	chars := strings.Split(fdisk, "")
	for i, c := range chars {
		if c == "." {
			continue
		}
		num := int(c[0] - '0') // Convert char to int
		checkSum += num * i
	}
	return checkSum
}
