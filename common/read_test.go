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
			size:     Point{1, 1},
		},
		"full": {
			input: `XY
#$`,
			expected: map[Point]rune{
				Point{0, 0}: 'X',
				Point{1, 0}: 'Y',
				Point{0, 1}: '#',
				Point{1, 1}: '$',
			},
			size: Point{1, 1},
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
