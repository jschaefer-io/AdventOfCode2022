package day04

import (
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

type area struct {
	from int
	to   int
}

type pair [2]area

func Solver() orchestration.Solver {
	return &day04{}
}

type day04 struct {
	pairs []pair
}

func (d *day04) Setup(data string) error {
	pairs := make([]pair, 0)
	for _, pairData := range strings.Split(data, "\n") {
		ranges := strings.Split(pairData, ",")

		lData := strings.Split(ranges[0], "-")
		lF, _ := strconv.Atoi(lData[0])
		lT, _ := strconv.Atoi(lData[1])

		rData := strings.Split(ranges[1], "-")
		rF, _ := strconv.Atoi(rData[0])
		rT, _ := strconv.Atoi(rData[1])
		pairs = append(pairs, pair{
			area{
				from: lF,
				to:   lT,
			},
			area{
				from: rF,
				to:   rT,
			},
		})
	}
	d.pairs = pairs
	return nil
}

func (d *day04) A() (string, error) {
	count := 0
	for _, p := range d.pairs {
		if (p[0].from <= p[1].from && p[0].to >= p[1].to) || (p[1].from <= p[0].from && p[1].to >= p[0].to) {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func (d *day04) B() (string, error) {
	count := 0
	for _, p := range d.pairs {
		if (p[0].from >= p[1].from && p[0].from <= p[1].to) || (p[0].to >= p[1].from && p[0].to <= p[1].to) ||
			(p[1].from >= p[0].from && p[1].from <= p[0].to) || (p[1].to >= p[0].from && p[1].to <= p[0].to) {
			count++
		}
	}
	return strconv.Itoa(count), nil
}
