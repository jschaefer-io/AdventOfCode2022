package day02

import (
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day02{}
}

type result struct {
	a rune
	b rune
}

type day02 struct {
	results []result
}

func (d *day02) Setup(data string) error {
	results := make([]result, 0)
	for _, match := range strings.Split(data, "\n") {
		res := strings.Split(match, " ")
		results = append(results, result{
			a: rune(res[0][0]),
			b: rune(res[1][0]),
		})
	}
	d.results = results
	return nil
}

func (d *day02) A() (string, error) {
	points := 0
	for _, match := range d.results {
		a := translationMap[match.a]
		b := translationMap[match.b]
		points += resolveGame(b, a) + values[b]
	}
	return strconv.Itoa(points), nil
}

func (d *day02) B() (string, error) {
	points := 0
	for _, match := range d.results {
		a := translationMap[match.a]
		b := ruleSet[a][match.b-'X']
		points += resolveGame(b, a) + values[b]
	}
	return strconv.Itoa(points), nil
}
