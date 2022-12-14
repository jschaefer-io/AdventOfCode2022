package day14

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day14{}
}

type position struct {
	x int
	y int
}

type field struct {
	minX *int
	maxX *int
	minY *int
	maxY *int
	area map[position]int
}

func (f *field) copy() *field {
	copyField := &field{
		minX: f.minX,
		maxX: f.maxX,
		minY: f.minY,
		maxY: f.maxY,
		area: make(map[position]int),
	}
	for pos, v := range f.area {
		copyField.area[pos] = v
	}
	return copyField
}

func (f *field) checkField(pos position) (bool, error) {
	if pos.x < *f.minX || pos.x > *f.maxX || pos.y > *f.maxY {
		return false, fmt.Errorf("field out of bounds")
	}
	nextField := f.area[pos]
	return nextField == 0, nil
}

func (f *field) spawnSand(origin position) (bool, position) {
	currentPos := origin
LOOP:
	for {
		tests := []position{
			{currentPos.x, currentPos.y + 1},
			{currentPos.x - 1, currentPos.y + 1},
			{currentPos.x + 1, currentPos.y + 1},
		}
		for _, test := range tests {
			available, err := f.checkField(test)
			if err != nil {
				return false, test
			}
			if available {
				currentPos = test
				continue LOOP
			}
		}
		f.set(currentPos, 2)
		return true, currentPos
	}
}

func (f *field) set(pos position, val int) {
	if f.minX == nil || *f.minX > pos.x {
		f.minX = &pos.x
	}
	if f.maxX == nil || *f.maxX < pos.x {
		f.maxX = &pos.x
	}
	if f.minY == nil || *f.minY > pos.y {
		f.minY = &pos.y

	}
	if f.maxY == nil || *f.maxY < pos.y {
		f.maxY = &pos.y
	}
	f.area[pos] = val
}

func (f *field) String() string {
	if len(f.area) == 0 {
		return ""
	}
	str := strings.Builder{}
	for y := *f.minY; y <= *f.maxY; y++ {
		for x := *f.minX; x <= *f.maxX; x++ {
			point := f.area[position{x, y}]
			switch point {
			case 1:
				str.WriteRune('#')
			case 2:
				str.WriteRune('o')
			default:
				str.WriteRune(' ')
			}
		}
		str.WriteRune('\n')
	}
	return str.String()
}

type day14 struct {
	field *field
}

func (d *day14) Setup(data string) error {
	d.field = &field{
		area: make(map[position]int),
	}
	data = strings.Trim(data, "\n")
	for _, line := range strings.Split(data, "\n") {
		rawPoints := strings.Split(line, " -> ")
		points := make([]position, len(rawPoints))
		for i, rawPoint := range rawPoints {
			split := strings.Split(rawPoint, ",")
			x, err := strconv.Atoi(split[0])
			if err != nil {
				return err
			}
			y, err := strconv.Atoi(split[1])
			if err != nil {
				return err
			}
			points[i] = position{x, y}
		}
		if len(points) <= 1 {
			return fmt.Errorf("no points provided")
		}
		current := points[0]
		next := points[1:]
		for len(next) > 0 {
			xLower := int(math.Min(float64(current.x), float64(next[0].x)))
			xUpper := int(math.Max(float64(current.x), float64(next[0].x)))
			yLower := int(math.Min(float64(current.y), float64(next[0].y)))
			yUpper := int(math.Max(float64(current.y), float64(next[0].y)))
			for x := xLower; x <= xUpper; x++ {
				for y := yLower; y <= yUpper; y++ {
					d.field.set(position{x, y}, 1)
				}
			}
			current = next[0]
			next = next[1:]
		}
	}
	return nil
}

func (d *day14) A() (string, error) {
	fieldInstance := d.field.copy()
	count := 0
	for {
		if rest, _ := fieldInstance.spawnSand(position{500, 0}); !rest {
			break
		}
		count++
	}
	return strconv.Itoa(count), nil
}

func (d *day14) B() (string, error) {
	fieldInstance := d.field.copy()
	ground := *d.field.maxY + 2

	for x := *d.field.minX - 1000; x <= *d.field.maxX+1000; x++ {
		fieldInstance.set(position{x, ground}, 1)
	}

	origin := position{500, 0}
	count := 0
	for {
		rest, lastPos := fieldInstance.spawnSand(origin)
		if !rest {
			return "", fmt.Errorf("void reached")
		}
		count++
		if lastPos == origin {
			break
		}
	}
	return strconv.Itoa(count), nil
}
