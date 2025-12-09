package main

import (
	"log/slog"
	"os"
	"time"

	"aoc2025/internal/puzzle7"
)

const inputPath = `/tmp/aoc/puzzle7`

func main() {
	start := time.Now()

	x, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	res, err := puzzle7.Run(x)
	if err != nil {
		panic(err)
	}

	slog.Info("success", "result", res, "duration_ms", time.Since(start).Milliseconds())
}
