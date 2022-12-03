package main

import (
	"fmt"

	"github.com/jschaefer-io/AdventOfCode2022/day01"
	"github.com/jschaefer-io/AdventOfCode2022/day02"
	"github.com/jschaefer-io/AdventOfCode2022/day03"
	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func main() {
	days := map[int]orchestration.Solver{
		1: day01.Solver(),
		2: day02.Solver(),
		3: day03.Solver(),
	}
	results := orchestration.Dispatch(days, "./inputs", 1, 3)
	for key, result := range results {
		if result.Error != nil {
			fmt.Println(fmt.Sprintf("Day %02d: %v", key, result.Error))
		} else {
			fmt.Printf("Day %02d (%s):\nA (%s): %s\nB (%s): %s\n", key, result.Time, result.TimeA, result.A, result.TimeB, result.B)
		}
		fmt.Println()
	}
}
