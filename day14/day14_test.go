package day14

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"
	"github.com/stretchr/testify/assert"
)

func TestApplyMask(t *testing.T) {
	tests := map[string]struct {
		value    int
		mask     string
		expected int
	}{
		"ex1": {
			value:    11,
			mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			expected: 73,
		},
		"ex2": {
			value:    101,
			mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			expected: 101,
		},
		"ex3": {
			value:    0,
			mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			expected: 64,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			and, or := masks(test.mask)
			actual := apply(test.value, and, or)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestExample1(t *testing.T) {
	actual := processInstructions(strings.NewReader(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`))
	assert.Equal(t, 165, actual)
}

func TestPart1(t *testing.T) {
	actual := processInstructions(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 4297467072083, actual)
}

func TestExample2(t *testing.T) {
	actual := processInstructions2(strings.NewReader(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`))
	assert.Equal(t, 208, actual)
}

func TestPart2(t *testing.T) {
	actual := processInstructions2(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 5030603328768, actual)
}
