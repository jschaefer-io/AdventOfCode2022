package day01

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day01{}
}

type day01 struct {
	list []int
}

func (d *day01) Setup(data string) error {
	list := make([]int, 0)
	groups := strings.Split(data, "\n\n")
	for _, group := range groups {
		count := 0
		for _, item := range strings.Split(group, "\n") {
			v, err := strconv.Atoi(item)
			if err != nil {
				return fmt.Errorf("unable to parse item as int: %w", err)
			}
			count += v
		}
		list = append(list, count)
	}
	if len(list) < 3 {
		return errors.New("less than 3 elfs in input set")
	}
	d.list = list
	sort.Ints(d.list)
	return nil
}

func (d *day01) A() (string, error) {
	return strconv.Itoa(d.list[len(d.list)-1]), nil
}

func (d *day01) B() (string, error) {
	sum := 0
	for _, c := range d.list[len(d.list)-3:] {
		sum += c
	}
	return strconv.Itoa(sum), nil
}
