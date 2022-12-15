package day15

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day15{}
}

type position struct {
	x int
	y int
}

func (p position) distance(pos position) int {
	return int(math.Abs(float64(p.x-pos.x)) + math.Abs(float64(p.y-pos.y)))
}

type day15 struct {
	area    map[position]rune
	pairs   map[position]position
	beacons []position
	sensors []position
}

func (d *day15) Setup(data string) error {
	d.pairs = make(map[position]position)
	d.area = make(map[position]rune)
	d.beacons = make([]position, 0)
	d.sensors = make([]position, 0)
	exp := regexp.MustCompile("(|-)\\d+")
	for _, line := range strings.Split(data, "\n") {
		nums := exp.FindAllString(line, -1)
		xS, _ := strconv.Atoi(nums[0])
		yS, _ := strconv.Atoi(nums[1])
		d.area[position{xS, yS}] = 'S'
		d.sensors = append(d.sensors, position{xS, yS})

		xB, _ := strconv.Atoi(nums[2])
		yB, _ := strconv.Atoi(nums[3])
		d.area[position{xB, yB}] = 'B'
		d.beacons = append(d.beacons, position{xB, yB})

		d.pairs[position{xS, yS}] = position{xB, yB}
	}
	return nil
}

func (d *day15) A() (string, error) {
	yPos := 10
	check := make(map[position]struct{})
	for sensor, beacon := range d.pairs {
		distance := sensor.distance(beacon)
		yDistance := sensor.distance(position{sensor.x, yPos})
		offset := distance - yDistance
		if offset >= 0 {
			for xO := offset * -1; xO <= offset; xO++ {
				pos := position{sensor.x + xO, yPos}
				if _, ok := check[pos]; !ok {
					if v := d.area[pos]; v != 'B' && v != 'S' {
						check[pos] = struct{}{}
					}
				}
			}
		}
	}
	return strconv.Itoa(len(check)), nil
}

func (d *day15) B() (string, error) {
	rangeOffset := 4000000
	var pos position
	var distance int
	var checkDistance int
	var found bool
	for y := 0; y <= rangeOffset; y++ {
		for x := 0; x <= rangeOffset; x++ {
			pos = position{x, y}
			found = false
			for sensor, beacon := range d.pairs {
				distance = sensor.distance(beacon)
				checkDistance = sensor.distance(pos)
				if checkDistance <= distance {
					x += distance - checkDistance
					found = true
					break
				}
			}
			if !found {
				return strconv.Itoa(pos.x*4000000 + pos.y), nil
			}
		}
	}
	return "", fmt.Errorf("distress beacon not found")
}
