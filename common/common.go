package common

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jpillora/puzzler/harness/aoc/user"
)

func ReadInputExample() string {
	wd, _ := os.Getwd()
	filepath := filepath.Join(wd, "input-example.txt")
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func RunDev(run user.RunFn) {
	input := ReadInputExample()
	fmt.Println(run(false, input))
}
