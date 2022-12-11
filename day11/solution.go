package day11

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

func Solver() orchestration.Solver {
	return &day11{}
}

type monkey struct {
	items          []uint64
	operation      []string
	test           uint64
	conditionTrue  int
	conditionFalse int
}

func (m *monkey) inspectItem(monkeys []*monkey, tureRes func(uint64) uint64) error {
	// Get next item
	item := m.items[0]
	m.items = m.items[1:]

	// Execute Operation
	var res uint64 = 0
	ai, err := strconv.Atoi(m.operation[0])
	a := uint64(ai)
	if err != nil {
		a = item
	}
	bi, err := strconv.Atoi(m.operation[2])
	b := uint64(bi)
	if err != nil {
		b = item
	}
	switch m.operation[1] {
	case "+":
		res = a + b
	case "*":
		res = a * b
	default:
		return fmt.Errorf("undefined operation %s", m.operation[1])
	}

	// lazy handler
	trueRes := tureRes(res)

	// find target monkey
	target := m.conditionFalse
	if trueRes%m.test == 0 {
		target = m.conditionTrue
	}
	monkeys[target].items = append(monkeys[target].items, trueRes)
	return nil
}

type day11 struct {
	monkeys []*monkey
	items   map[int][]uint64
}

func (d *day11) Setup(data string) error {
	d.items = make(map[int][]uint64)
	monkeys := make([]*monkey, 0)
	exp := regexp.MustCompile("\\d+")
	opExp := regexp.MustCompile(".*?=")
	for id, mky := range strings.Split(data, "\n\n") {
		m := &monkey{
			items: make([]uint64, 0),
		}
		lines := strings.Split(mky, "\n")

		// items
		items := exp.FindAllString(lines[1], -1)
		mItems := make([]uint64, 0)
		for _, item := range items {
			v, err := strconv.Atoi(item)
			if err != nil {
				return err
			}
			mItems = append(mItems, uint64(v))
		}
		d.items[id] = mItems

		// test
		test, err := strconv.Atoi(exp.FindString(lines[3]))
		if err != nil {
			return err
		}
		m.test = uint64(test)

		// conditions
		a, err := strconv.Atoi(exp.FindString(lines[4]))
		if err != nil {
			return err
		}
		m.conditionTrue = a
		b, err := strconv.Atoi(exp.FindString(lines[5]))
		if err != nil {
			return err
		}
		m.conditionFalse = b

		m.operation = strings.Split(strings.Trim(opExp.ReplaceAllString(lines[2], ""), " "), " ")
		monkeys = append(monkeys, m)
	}
	d.monkeys = monkeys
	return nil
}

func (d *day11) countActiveApes(rounds int, handler func(uint64) uint64) (int, error) {
	for i, m := range d.monkeys {
		m.items = append(make([]uint64, 0), d.items[i]...)
	}
	counts := make(map[int]int)
	var err error
	for i := 0; i < rounds; i++ {
		for mId, m := range d.monkeys {
			for len(m.items) > 0 {
				counts[mId]++
				err = m.inspectItem(d.monkeys, handler)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	list := make([]int, len(d.monkeys))
	for i, _ := range d.monkeys {
		list[i] = counts[i]
	}
	sort.Ints(list)
	return list[len(d.monkeys)-2] * list[len(d.monkeys)-1], nil
}

func (d *day11) A() (string, error) {
	res, err := d.countActiveApes(20, func(res uint64) uint64 {
		return res / 3
	})
	return strconv.Itoa(res), err
}

func (d *day11) B() (string, error) {
	// Find the LCM, since we get overflow the stress levels
	testNums := make([]uint64, len(d.monkeys))
	for i, m := range d.monkeys {
		testNums[i] = m.test
	}
	lowestCommonMultiple := lcm(testNums[0], testNums[1], testNums[2:]...)

	res, err := d.countActiveApes(10000, func(res uint64) uint64 {
		return res % lowestCommonMultiple
	})
	return strconv.Itoa(res), err
}
