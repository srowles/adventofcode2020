package day7

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"

	"github.com/stretchr/testify/assert"
)

func TestLineParse(t *testing.T) {
	tests := map[string]struct {
		input  string
		outer  bag
		inners []bag
	}{
		"nothing": {
			outer: bag{},
		},
		"singleInner": {
			input: "bright white bags contain 1 shiny gold bag.",
			outer: bag{colour: "bright white", number: 1},
			inners: []bag{
				bag{colour: "shiny gold", number: 1},
			},
		},
		"twoInner": {
			input: "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			outer: bag{colour: "light red", number: 1},
			inners: []bag{
				bag{colour: "bright white", number: 1},
				bag{colour: "muted yellow", number: 2},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			bag, bags := parseLine(test.input)
			assert.Equal(t, test.outer, bag)
			assert.Equal(t, test.inners, bags)
		})
	}
}

func TestExample(t *testing.T) {
	actual := process(strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`))
	assert.Equal(t, 4, actual)
}

func TestPart1(t *testing.T) {
	actual := process(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 124, actual)
}

func TestExample2(t *testing.T) {
	actual := process2(strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`))
	assert.Equal(t, 32, actual)
}

func TestPart2(t *testing.T) {
	actual := process2(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 34862, actual)
}
