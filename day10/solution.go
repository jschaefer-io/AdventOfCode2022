package day10

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day10{}
}

type instruction struct {
	method   string
	argument int
}

type day10 struct {
	instructions []instruction
}

func (d *day10) Setup(data string) error {
	d.instructions = make([]instruction, 0)
	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, " ")
		v := 0
		if len(parts) == 2 {
			parsed, err := strconv.Atoi(parts[1])
			if err != nil {
				return err
			}
			v = parsed
		}
		d.instructions = append(d.instructions, instruction{
			method:   parts[0],
			argument: v,
		})

	}
	return nil
}

func (d *day10) runWith(handler func(cycle, x int)) error {
	x := 1
	cycle := 0
	for _, ins := range d.instructions {
		switch ins.method {
		case "noop":
			cycle++
			handler(cycle, x)
		case "addx":
			cycle++
			handler(cycle, x)
			cycle++
			handler(cycle, x)
			x += ins.argument
		default:
			return fmt.Errorf("undefined operation %s", ins.method)
		}
	}
	return nil
}

func (d *day10) A() (string, error) {
	sum := 0
	err := d.runWith(func(cycle, x int) {
		if cycle == 20 || (cycle > 20 && (cycle-20)%40 == 0) {
			sum += x * cycle
		}
	})
	return strconv.Itoa(sum), err
}

func (d *day10) B() (string, error) {
	var screen [6][40]rune
	err := d.runWith(func(cycle, x int) {
		pos := cycle - 1
		row := pos / 40
		pixel := pos % 40

		if math.Abs(float64(pixel-x)) <= 1 {
			screen[row][pixel] = '■'
		} else {
			screen[row][pixel] = ' '
		}
	})
	str := strings.Builder{}
	str.WriteRune('\n')
	for _, line := range screen {
		for _, char := range line {
			str.WriteRune(char)
		}
		str.WriteRune('\n')
	}
	return str.String(), err
}
