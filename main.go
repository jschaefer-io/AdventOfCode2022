package main

import (
	"fmt"
	"log"

	"github.com/jschaefer-io/AdventOfCode2022/day01"
	"github.com/jschaefer-io/AdventOfCode2022/day02"
	"github.com/jschaefer-io/AdventOfCode2022/day03"
	"github.com/jschaefer-io/AdventOfCode2022/day04"
	"github.com/jschaefer-io/AdventOfCode2022/day05"
	"github.com/jschaefer-io/AdventOfCode2022/day06"
	"github.com/jschaefer-io/AdventOfCode2022/day07"
	"github.com/jschaefer-io/AdventOfCode2022/day08"
	"github.com/jschaefer-io/AdventOfCode2022/day09"
	"github.com/jschaefer-io/AdventOfCode2022/day10"
	"github.com/jschaefer-io/AdventOfCode2022/day11"
	"github.com/jschaefer-io/AdventOfCode2022/day12"
	"github.com/jschaefer-io/AdventOfCode2022/day13"
	"github.com/jschaefer-io/AdventOfCode2022/day14"
	"github.com/jschaefer-io/AdventOfCode2022/day15"
	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func main() {
	days := []orchestration.Solver{
		day01.Solver(),
		day02.Solver(),
		day03.Solver(),
		day04.Solver(),
		day05.Solver(),
		day06.Solver(),
		day07.Solver(),
		day08.Solver(),
		day09.Solver(),
		day10.Solver(),
		day11.Solver(),
		day12.Solver(),
		day13.Solver(),
		day14.Solver(),
		day15.Solver(),
	}
	results, err := orchestration.Dispatch(days, "./inputs", 15, 15)
	if err != nil {
		log.Fatalln(err)
	}
	for key, result := range results {
		if result.Error != nil {
			fmt.Println(fmt.Sprintf("Day %02d: %v", key, result.Error))
		} else {
			fmt.Printf("Day %02d (%s):\nA (%s): %s\nB (%s): %s\n", key, result.Time, result.TimeA, result.A, result.TimeB, result.B)
		}
		fmt.Println()
	}
}
