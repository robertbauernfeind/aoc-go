package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

	filepath3 := filepath.Join(wd, "input-user.txt")
	_, err = os.Stat(filepath3)
	if err != nil {
		return
	}

	input3 := readInputExample(filepath3)
	fmt.Println("Part1: ", run(false, input3))
}

func SplitToLines(s string) []string {
	return strings.Split(s, "\n")
}

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return v
}

func SliceAtoi(s []string) []int {
	nums := []int{}
	for _, v := range s {
		nums = append(nums, Atoi(v))
	}
	return nums
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}
