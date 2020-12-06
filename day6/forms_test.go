package day6

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/srowles/adventofcode2020/common"
)

func TestExample(t *testing.T) {
	actual, _ := process(strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b
`))
	assert.Equal(t, 11, actual)
}

func TestPart1(t *testing.T) {
	actual, _ := process(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 6585, actual)
}

func TestExample2(t *testing.T) {
	_, actual := process(strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b
`))
	assert.Equal(t, 6, actual)
}

func TestPart2(t *testing.T) {
	_, actual := process(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 3276, actual)
}
