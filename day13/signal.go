package day13

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type signal interface {
	fmt.Stringer
	isList() bool
	toList() signal
	compare(signal) (bool, bool)
	value() int
}

func parseSignal(value string) (signal, error) {
	items := make([]any, 0)
	err := json.Unmarshal([]byte(value), &items)
	if err != nil {
		return nil, err
	}
	var parseItems func(any) (signal, error)
	parseItems = func(item any) (signal, error) {
		switch v := item.(type) {
		case float64:
			return valueSignal(int(v)), nil
		case []any:
			sig := listSignal{
				items: make([]signal, len(v)),
			}
			for i, s := range v {
				nv, err := parseItems(s)
				if err != nil {
					return nil, err
				}
				sig.items[i] = nv
			}
			return sig, nil
		default:
			return nil, fmt.Errorf("unknown type for %v", v)
		}
	}
	return parseItems(items)
}

type valueSignal int

func (v valueSignal) String() string {
	return strconv.Itoa(v.value())
}

func (v valueSignal) isList() bool {
	return false
}

func (v valueSignal) toList() signal {
	return listSignal{
		items: []signal{v},
	}
}

func (v valueSignal) compare(s signal) (bool, bool) {
	if s.isList() {
		return v.toList().compare(s)
	} else {
		return v.value() < s.value(), v.value() == s.value()
	}
}

func (v valueSignal) value() int {
	return int(v)
}

// --

type listSignal struct {
	items []signal
}

func (l listSignal) String() string {
	str := strings.Builder{}
	str.WriteRune('[')
	for i, s := range l.items {
		if i != 0 {
			str.WriteRune(',')
		}
		str.WriteString(s.String())
	}
	str.WriteRune(']')
	return str.String()
}

func (l listSignal) value() int {
	return l.items[0].value()
}

func (l listSignal) isList() bool {
	return true
}

func (l listSignal) toList() signal {
	return l
}

func (l listSignal) compare(s signal) (bool, bool) {
	if !s.isList() {
		return l.compare(s.toList())
	}
	compare := s.(listSignal)
	lCount := len(l.items)
	rCount := len(compare.items)
	i := 0
	for i < lCount || i < rCount {
		if i >= lCount {
			return true, false
		} else if i >= rCount {
			return false, false
		}
		if res, same := l.items[i].compare(compare.items[i]); !same {
			return res, false
		}
		i++
	}
	return true, true
}
