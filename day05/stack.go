package day05

import (
	"fmt"
)

type stack struct {
	items []rune
}

func (s *stack) push(item rune) {
	s.items = append(s.items, item)
}

func (s *stack) add(batch *stack) {
	for {
		v, err := batch.pop()
		if err != nil {
			return
		}
		s.push(v)
	}
}

func (s *stack) copy() *stack {
	count := len(s.items)
	newStack := &stack{items: make([]rune, count)}
	for i := 0; i < count; i++ {
		newStack.items[i] = s.items[i]
	}
	return newStack
}

func (s *stack) peek() (rune, error) {
	count := len(s.items)
	if count == 0 {
		return ' ', fmt.Errorf("stack empty")
	}
	return s.items[count-1], nil
}

func (s *stack) pop() (rune, error) {
	v, err := s.peek()
	if err != nil {
		return v, err
	}
	s.items = s.items[:len(s.items)-1]
	return v, nil
}
