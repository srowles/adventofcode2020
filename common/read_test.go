package common

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapRead(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected map[Point]rune
		size     Point
	}{
		"empty": {
			expected: map[Point]rune{},
		},
		"emptyWithEmptyInput": {
			input: `..
..`,
			expected: map[Point]rune{},
			size:     Point{X: 1, Y: 1, Z: 0, W: 0},
		},
		"full": {
			input: `XY
#$`,
			expected: map[Point]rune{
				Point{X: 0, Y: 0, Z: 0, W: 0}: 'X',
				Point{X: 1, Y: 0, Z: 0, W: 0}: 'Y',
				Point{X: 0, Y: 1, Z: 0, W: 0}: '#',
				Point{X: 1, Y: 1, Z: 0, W: 0}: '$',
			},
			size: Point{1, 1, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, point := ReadMap(strings.NewReader(test.input))
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.size, point)
		})
	}
}
