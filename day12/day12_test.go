package day12

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	actual := processInstructions(strings.NewReader(`F10
N3
F7
R90
F11`))
	assert.Equal(t, 25, actual)
}

func TestDirChange(t *testing.T) {
	tests := map[string]struct {
		instruction string
		val         int
		current     direction
		expected    direction
	}{
		"none": {},
		"eastR90": {
			instruction: "R",
			val:         90,
			current:     east,
			expected:    south,
		},
		"westR90": {
			instruction: "R",
			val:         90,
			current:     west,
			expected:    north,
		},
		"westR270": {
			instruction: "R",
			val:         270,
			current:     west,
			expected:    south,
		},
		"eastL90": {
			instruction: "L",
			val:         90,
			current:     east,
			expected:    north,
		},
		"eastL270": {
			instruction: "L",
			val:         270,
			current:     east,
			expected:    south,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := processDirChange(test.instruction, test.val, test.current)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestPart1(t *testing.T) {
	actual := processInstructions(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 796, actual)
}

func TestExample2(t *testing.T) {
	actual := processInstructions2(strings.NewReader(`F10
N3
F7
R90
F11`))
	assert.Equal(t, 286, actual)
}

func TestPart2(t *testing.T) {
	actual := processInstructions2(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 39446, actual)
}
