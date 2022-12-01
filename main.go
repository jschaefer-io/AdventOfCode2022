package main

import (
	"fmt"

	"github.com/jschaefer-io/AdventOfCode2022/day01"
	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func main() {
	days := map[int]orchestration.Solver{
		1: day01.Solver(),
	}
	results := orchestration.Dispatch(days, "./inputs", 1, 1)
	for key, result := range results {
		if result.Error != nil {
			fmt.Println(fmt.Sprintf("Day %02d: %v", key, result.Error))
		} else {
			fmt.Printf("Day %02d (%s):\nA (%s): %s\nB (%s): %s\nErr: %v", key, result.Time, result.TimeA, result.A, result.TimeB, result.B, result.Error)
		}
		fmt.Println()
	}
}
