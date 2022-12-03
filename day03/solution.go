package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

type item rune

func (i item) priority() int {
	if i <= 'Z' {
		return int(i - 'A' + 27)
	} else {
		return int(i - 'a' + 1)
	}
}

func Solver() orchestration.Solver {
	return &day03{}
}

type day03 struct {
	packs [][]item
}

func (d *day03) Setup(data string) error {
	backpacks := make([][]item, 0)
	for _, pack := range strings.Split(data, "\n") {
		backpacks = append(backpacks, []item(pack))
	}
	d.packs = backpacks
	return nil
}

func (d *day03) A() (string, error) {
	sum := 0
LOOP:
	for _, pack := range d.packs {
		count := len(pack)
		left := make(map[item]struct{})
		for _, it := range pack[:count/2] {
			left[it] = struct{}{}
		}
		for _, it := range pack[count/2:] {
			if _, ok := left[it]; ok {
				sum += it.priority()
				continue LOOP
			}
		}
		return "", fmt.Errorf("no double item found in pack %s", string(pack))
	}
	return strconv.Itoa(sum), nil
}

func (d *day03) B() (string, error) {
	sum := 0
LOOP:
	for i := 0; i < len(d.packs); i += 3 {
		group := d.packs[i : i+3]
		check := make(map[item]int)
		for _, elf := range group {
			elfCheck := make(map[item]struct{})
			for _, it := range elf {
				if _, ok := elfCheck[it]; ok {
					continue
				}
				check[it]++
				elfCheck[it] = struct{}{}
			}
		}
		for it, count := range check {
			if count == 3 {
				sum += it.priority()
				continue LOOP
			}
		}
		return "", fmt.Errorf("group %d has no common item", i/3)
	}
	return strconv.Itoa(sum), nil
}
