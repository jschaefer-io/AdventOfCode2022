package day06

import (
	"fmt"
	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
	"strconv"
	"strings"
)

func Solver() orchestration.Solver {
	return &day06{}
}

type day06 struct {
	message string
}

func (d *day06) Setup(data string) error {
	str := strings.Trim(data, "\n")
	if len(str) < 4 {
		return fmt.Errorf("unsufficient message length")
	}
	d.message = str
	return nil
}

func (d *day06) findMarkerPosition(charCount int) (int, error) {
LOOP:
	for i := 3; i < len(d.message); i++ {
		list := make(map[rune]struct{})
		for u := 0; u < charCount; u++ {
			if _, ok := list[rune(d.message[i-u])]; ok {
				continue LOOP
			}
			list[rune(d.message[i-u])] = struct{}{}
		}
		return i + 1, nil
	}
	return 0, fmt.Errorf("no marker found")
}

func (d *day06) A() (string, error) {
	v, err := d.findMarkerPosition(4)
	return strconv.Itoa(v), err
}

func (d *day06) B() (string, error) {
	v, err := d.findMarkerPosition(14)
	return strconv.Itoa(v), err
}
