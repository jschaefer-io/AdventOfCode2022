package day13

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day13{}
}

type day13 struct {
	pairs [][2]signal
}

func (d *day13) Setup(data string) error {
	d.pairs = make([][2]signal, 0)
	for _, rawPair := range strings.Split(data, "\n\n") {
		pair := strings.Split(rawPair, "\n")
		left, err := parseSignal(pair[0])
		if err != nil {
			return err
		}
		right, err := parseSignal(pair[1])
		if err != nil {
			return err
		}
		d.pairs = append(d.pairs, [2]signal{
			left,
			right,
		})
	}
	return nil
}

func (d *day13) A() (string, error) {
	sum := 0
	for i, pair := range d.pairs {
		order, undefined := pair[0].compare(pair[1])
		if undefined {
			return "", fmt.Errorf("unable to compare pair: %s %s", pair[0], pair[1])
		}
		if order {
			sum += i + 1
		}
	}
	return strconv.Itoa(sum), nil
}

func (d *day13) B() (string, error) {
	list := make([]signal, 0)
	first, err := parseSignal("[[2]]")
	if err != nil {
		return "", err
	}
	second, err := parseSignal("[[6]]")
	if err != nil {
		return "", err
	}
	list = append(list, first, second)
	for _, pair := range d.pairs {
		list = append(list, pair[0], pair[1])
	}
	sort.Slice(list, func(i, j int) bool {
		sortRes, _ := list[i].compare(list[j])
		return sortRes
	})

	var firstIndex int
	var secondIndex int
	for i, item := range list {
		if item.String() == first.String() {
			firstIndex = i + 1
		} else if item.String() == second.String() {
			secondIndex = i + 1
		}
	}

	if firstIndex == 0 || secondIndex == 0 {
		return "", fmt.Errorf("reference signals not found")
	}

	return strconv.Itoa(firstIndex * secondIndex), nil
}
