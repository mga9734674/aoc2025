package main

import (
	"aoc2025/internal/puzzle3"
	"fmt"
)

const inputPath = `/tmp/aoc/puzzle3`

func main() {
	res, err := puzzle3.ParseAndRun(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", res)
}
