package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/srowles/adventofcode2020/common"
)

func TestPart1(t *testing.T) {
	result := part1(common.ReaderFromInts([]int{1721, 979, 366, 299, 675, 1456}))
	assert.Equal(t, 514579, result)
}

func TestRunPart1(t *testing.T) {
	result := part1(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 786811, result)
}

func TestRunPart2(t *testing.T) {
	result := part2(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 199068980, result)
}
