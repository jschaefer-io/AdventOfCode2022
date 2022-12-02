package day02

type valueType uint8

const (
	rock valueType = iota
	paper
	scissors
)

var values = map[valueType]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

var translationMap = map[rune]valueType{
	'A': rock,
	'B': paper,
	'C': scissors,
	'X': rock,
	'Y': paper,
	'Z': scissors,
}

var ruleSet = map[valueType][3]valueType{
	rock:     {scissors, rock, paper},
	paper:    {rock, paper, scissors},
	scissors: {paper, scissors, rock},
}

func resolveGame(player valueType, opponent valueType) int {
	rules := ruleSet[player]
	switch opponent {
	case rules[0]:
		return 6
	case rules[1]:
		return 3
	default:
		return 0
	}
}
