package common

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jpillora/puzzler/harness/aoc/user"
)

func readInputExample(file string) string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func RunDev(run user.RunFn) {
	wd, _ := os.Getwd()
	filepath1 := filepath.Join(wd, "input-example.txt")
	input1 := readInputExample(filepath1)
	fmt.Println("Part1: ", run(false, input1))

	filepath2 := filepath.Join(wd, "input-example2.txt")
	_, err := os.Stat(filepath2)
	input2 := input1
	if err == nil {
		input2 = readInputExample(filepath2)
	}

	fmt.Println("Part2: ", run(true, input2))
}
