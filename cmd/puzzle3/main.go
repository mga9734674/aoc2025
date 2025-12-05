package main

import (
	"fmt"

	"aoc2025/internal/puzzle3"
)

const inputPath = `/tmp/aoc/puzzle3`

func main() {
	res, err := puzzle3.ParseAndRun(inputPath, 12)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", res)
}
