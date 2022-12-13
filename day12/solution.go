package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

type position struct {
	x int
	y int
}

func Solver() orchestration.Solver {
	return &day12{}
}

type day12 struct {
	player position
	target position
	grid   [][]int
	width  int
	height int
}

func (d *day12) Setup(data string) error {
	foundPlayer := false
	foundTarget := false
	lines := strings.Split(data, "\n")
	grid := make([][]int, len(lines))
	for y, line := range lines {
		row := make([]int, len(line))
		for x, c := range line {
			char := c - 'a'
			if c == 'S' {
				char = 'a' - 'a'
				d.player = position{x, y}
				foundPlayer = true
			} else if c == 'E' {
				char = 'z' - 'a'
				d.target = position{x, y}
				foundTarget = true
			}
			row[x] = int(char)
		}
		grid[y] = row
	}
	d.width = len(grid[0])
	d.height = len(grid)
	d.grid = grid
	if !foundPlayer {
		return fmt.Errorf("no player position found")
	}
	if !foundTarget {
		return fmt.Errorf("no target position found")
	}
	return nil
}

func (d *day12) getPositions(pos position) []position {
	height := d.grid[pos.y][pos.x]
	list := make([]position, 0)
	offsets := []position{
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 0},
	}
	for _, offset := range offsets {
		if pos.x+offset.x >= 0 && pos.x+offset.x < d.width && pos.y+offset.y >= 0 && pos.y+offset.y < d.height {
			nextPos := position{
				pos.x + offset.x,
				pos.y + offset.y,
			}
			targetHeight := d.grid[nextPos.y][nextPos.x]
			if targetHeight-height > 1 {
				continue
			}
			list = append(list, nextPos)
		}
	}
	return list
}

// breadth first search
func (d *day12) findStepBetween(start, end position) (int, error) {
	stack := make([]position, 0)
	visited := make(map[position]int)
	var current position
	stack = append(stack, start)
	visited[current] = 0
	for len(stack) > 0 {
		current = stack[0]
		stack = stack[1:]
		for _, pos := range d.getPositions(current) {
			if _, ok := visited[pos]; ok {
				continue
			}
			visited[pos] = visited[current] + 1
			stack = append(stack, pos)
		}
	}
	steps, ok := visited[end]
	if !ok {
		return steps, fmt.Errorf("target not reachable")
	}
	return steps, nil
}

func (d *day12) A() (string, error) {
	res, err := d.findStepBetween(d.player, d.target)
	return strconv.Itoa(res), err
}

func (d *day12) B() (string, error) {
	min := -1
	for x := 0; x < d.width; x++ {
		for y := 0; y < d.height; y++ {
			if d.grid[y][x] != 0 {
				continue
			}
			check, err := d.findStepBetween(position{x, y}, d.target)
			if err == nil && (min == -1 || min > check) {
				min = check
			}
		}
	}
	if min == -1 {
		return "0", fmt.Errorf("no path found")
	}
	return strconv.Itoa(min), nil
}
