package day05

import (
	"strconv"
	"strings"

	"github.com/jschaefer-io/AdventOfCode2022/orchestration"
)

type instruction struct {
	count int
	from  int
	to    int
}

func Solver() orchestration.Solver {
	return &day05{}
}

type day05 struct {
	instructions []instruction
	stacks       []*stack
}

func (d *day05) getStacks() []*stack {
	list := make([]*stack, len(d.stacks))
	for i, s := range d.stacks {
		list[i] = s.copy()
	}
	return list
}

func (d *day05) Setup(data string) error {
	instructions := make([]instruction, 0)
	parts := strings.Split(data, "\n\n")

	stacks := make([]*stack, 0)
	stackPart := strings.Split(parts[0], "\n")
	for s := 1; s < len(stackPart[len(stackPart)-1]); s += 4 {
		nStack := &stack{
			items: make([]rune, 0),
		}
		for c := len(stackPart) - 2; c >= 0; c-- {
			if len(stackPart) <= c || len(stackPart[c]) <= s || stackPart[c][s] == ' ' {
				break
			}
			nStack.push(rune(stackPart[c][s]))
		}
		stacks = append(stacks, nStack)
	}

	for _, in := range strings.Split(parts[1], "\n") {
		vals := strings.Split(in, " ")
		count, _ := strconv.Atoi(vals[1])
		from, _ := strconv.Atoi(vals[3])
		to, _ := strconv.Atoi(vals[5])
		instructions = append(instructions, instruction{
			count: count,
			from:  from,
			to:    to,
		})
	}

	d.stacks = stacks
	d.instructions = instructions
	return nil
}

func (d *day05) executeWithStack(handler func(from *stack, to *stack, in instruction) error) (string, error) {
	stacks := d.getStacks()
	for _, in := range d.instructions {
		err := handler(stacks[in.from-1], stacks[in.to-1], in)
		if err != nil {
			return "", err
		}
	}
	str := strings.Builder{}
	for _, s := range stacks {
		top, err := s.peek()
		if err != nil {
			return "", err
		}
		str.WriteRune(top)
	}
	return str.String(), nil
}

func (d *day05) A() (string, error) {
	return d.executeWithStack(func(from *stack, to *stack, in instruction) error {
		for i := 0; i < in.count; i++ {
			val, err := from.pop()
			if err != nil {
				return err
			}
			to.push(val)
		}
		return nil
	})
}

func (d *day05) B() (string, error) {
	return d.executeWithStack(func(from *stack, to *stack, in instruction) error {
		crane := &stack{
			items: make([]rune, 0),
		}
		for i := 0; i < in.count; i++ {
			val, err := from.pop()
			if err != nil {
				return err
			}
			crane.push(val)
		}
		to.add(crane)
		return nil
	})
}
