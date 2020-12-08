package day8

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	actual := process(strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`))
	assert.Equal(t, 5, actual)
}

func TestPart1(t *testing.T) {
	actual := process(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 1816, actual)
}

func TestExample2(t *testing.T) {
	actual := process2(strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`))
	assert.Equal(t, 8, actual)
}

func TestPart2(t *testing.T) {
	actual := process2(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 1149, actual)
}
