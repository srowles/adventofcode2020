package day9

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	actual := findInvalidNumber(strings.NewReader(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`), 5)
	assert.Equal(t, 127, actual)
}

func TestPart1(t *testing.T) {
	actual := findInvalidNumber(common.ReaderFromFile("input.txt"), 25)
	assert.Equal(t, 1398413738, actual)
}

func TestExample2(t *testing.T) {
	actual := findSequence(strings.NewReader(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`), 127)
	assert.Equal(t, 62, actual)
}

func TestPart2(t *testing.T) {
	actual := findSequence(common.ReaderFromFile("input.txt"), 1398413738)
	assert.Equal(t, 169521051, actual)
}
