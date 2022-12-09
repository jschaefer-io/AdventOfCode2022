package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day09{}
}

type day09 struct {
	moves []coordinate
}

func (d *day09) Setup(data string) error {
	d.moves = make([]coordinate, 0)
	for line, move := range strings.Split(data, "\n") {
		v, err := strconv.Atoi(move[2:])
		if err != nil {
			return fmt.Errorf("unable to parse int %s on line %d", move[2:], line+1)
		}
		moveSet := coordinate{}
		switch move[0] {
		case 'L':
			moveSet.x = -1
		case 'R':
			moveSet.x = 1
		case 'U':
			moveSet.y = 1
		case 'D':
			moveSet.y = -1
		default:
			return fmt.Errorf("undefined move %s", string(move[0]))
		}
		for i := 0; i < v; i++ {
			d.moves = append(d.moves, moveSet)
		}
	}
	return nil
}

func (d *day09) A() (string, error) {
	tailVisits := make(map[coordinate]struct{})
	head := &coordinate{}
	tail := &coordinate{}
	tailVisits[*tail] = struct{}{}
	for _, move := range d.moves {
		head.x += move.x
		head.y += move.y
		tail.moveTowards(*head, 1)
		tailVisits[*tail] = struct{}{}
	}
	return strconv.Itoa(len(tailVisits)), nil
}

func (d *day09) B() (string, error) {
	tailVisits := make(map[coordinate]struct{})
	segmentCount := 10
	rope := make([]coordinate, segmentCount)
	tailVisits[rope[segmentCount-1]] = struct{}{}
	for _, move := range d.moves {
		rope[0].x += move.x
		rope[0].y += move.y
		for i := 1; i < segmentCount; i++ {
			rope[i].moveTowards(rope[i-1], 1)
		}
		tailVisits[rope[segmentCount-1]] = struct{}{}
	}
	return strconv.Itoa(len(tailVisits)), nil
}
