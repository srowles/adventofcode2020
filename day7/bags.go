package day7

import (
	"fmt"
	"io"
	"strings"

	"github.com/srowles/adventofcode2020/common"
)

type bag struct {
	colour string
	number int
}

type holder struct {
	bagMap map[string][]bag
}

func (h holder) count(name string, c int) int {
	bags := h.bagMap[name]
	if len(bags) == 0 {
		return 1
	}
	if len(bags) == 1 && bags[0].colour == "" {
		return 1
	}
	for _, b := range bags {
		c += b.number * h.count(b.colour, 0)
	}
	return c + 1
}

func (h holder) find(name string) bool {
	if name == "shiny gold" {
		return true
	}
	bags, ok := h.bagMap[name]
	if !ok {
		return false
	}
	for _, b := range bags {
		if h.find(b.colour) {
			return true
		}
	}

	return false
}

// dark indigo bags contain 1 light maroon bag, 3 pale red bags, 1 drab brown bag, 4 dim magenta bags.
func parseLine(line string) (bag, []bag) {
	if line == "" {
		return bag{}, nil
	}
	parts := strings.Split(line, " contain ")
	var a, b string
	fmt.Sscanf(parts[0], "%s %s bags", &a, &b)
	outer := bag{colour: strings.TrimSpace(fmt.Sprintf("%s %s", a, b)), number: 1}
	var inners []bag
	for _, description := range strings.Split(parts[1], ",") {
		var n int
		var a, b string
		fmt.Sscanf(description, "%d %s %s bag", &n, &a, &b)
		inners = append(inners, bag{colour: strings.TrimSpace(fmt.Sprintf("%s %s", a, b)), number: n})
	}
	return outer, inners
}

func process(reader io.Reader) int {
	contains := make(map[string][]bag)
	for _, line := range common.StringListFromReader(reader) {
		parent, bags := parseLine(line)
		if parent.colour == "" {
			continue
		}
		for _, b := range bags {
			list := contains[parent.colour]
			list = append(list, b)
			contains[parent.colour] = list
		}
	}

	count := 0
	h := holder{bagMap: contains}
	for edgeBag := range contains {
		if edgeBag == "shiny gold" {
			continue
		}
		if h.find(edgeBag) {
			count++
		}
	}

	return count
}

func process2(reader io.Reader) int {
	contains := make(map[string][]bag)
	for _, line := range common.StringListFromReader(reader) {
		parent, bags := parseLine(line)
		if parent.colour == "" {
			continue
		}
		for _, b := range bags {
			list := contains[parent.colour]
			list = append(list, b)
			contains[parent.colour] = list
		}
	}

	h := holder{bagMap: contains}
	count := h.count("shiny gold", 0)

	return count - 1
}
