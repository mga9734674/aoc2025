package main

import (
	"log/slog"
	"os"
	"time"

	"aoc2025/internal/puzzle8"
)

const inputPath = `/tmp/aoc/puzzle8`

func main() {
	start := time.Now()

	x, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	res, err := puzzle8.Run(x)
	if err != nil {
		panic(err)
	}

	slog.Info("success", "result", res, "duration_ms", time.Since(start).Milliseconds())
}
