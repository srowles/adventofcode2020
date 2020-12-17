package day17

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	actual := simulate(strings.NewReader(`.#.
..#
###
`), 6)
	assert.Equal(t, 112, actual)
}

func TestPart1(t *testing.T) {
	actual := simulate(common.ReaderFromFile("input.txt"), 6)
	assert.Equal(t, 348, actual)
}

func TestExample2(t *testing.T) {
	actual := simulate4(strings.NewReader(`.#.
..#
###
`), 6)
	assert.Equal(t, 848, actual)
}

func TestPart2(t *testing.T) {
	actual := simulate4(common.ReaderFromFile("input.txt"), 6)
	assert.Equal(t, 2236, actual)
}
