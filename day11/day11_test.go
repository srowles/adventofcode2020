package day11

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"
	"github.com/stretchr/testify/assert"
)

func TestRules(t *testing.T) {
	tests := map[string]struct {
		input    map[common.Point]rune
		expected map[common.Point]rune
	}{
		"allEmpty": {
			input: map[common.Point]rune{
				{X: -1, Y: -1}: 'L',
				{X: 0, Y: -1}:  'L',
				{X: 1, Y: -1}:  'L',
				{X: -1, Y: 0}:  'L',
				{X: 0, Y: 0}:   'L',
				{X: 1, Y: 0}:   'L',
				{X: -1, Y: 1}:  'L',
				{X: 0, Y: 1}:   'L',
				{X: 1, Y: 1}:   'L',
			},
			expected: map[common.Point]rune{
				{X: -1, Y: -1}: '#',
				{X: 0, Y: -1}:  '#',
				{X: 1, Y: -1}:  '#',
				{X: -1, Y: 0}:  '#',
				{X: 0, Y: 0}:   '#',
				{X: 1, Y: 0}:   '#',
				{X: -1, Y: 1}:  '#',
				{X: 0, Y: 1}:   '#',
				{X: 1, Y: 1}:   '#',
			},
		},
		"oneSeat": {
			input: map[common.Point]rune{
				{X: -1, Y: -1}: '.',
				{X: 0, Y: -1}:  '.',
				{X: 1, Y: -1}:  '.',
				{X: -1, Y: 0}:  '.',
				{X: 0, Y: 0}:   'L',
				{X: 1, Y: 0}:   '.',
				{X: -1, Y: 1}:  '.',
				{X: 0, Y: 1}:   '.',
				{X: 1, Y: 1}:   '.',
			},
			expected: map[common.Point]rune{
				{X: -1, Y: -1}: '.',
				{X: 0, Y: -1}:  '.',
				{X: 1, Y: -1}:  '.',
				{X: -1, Y: 0}:  '.',
				{X: 0, Y: 0}:   '#',
				{X: 1, Y: 0}:   '.',
				{X: -1, Y: 1}:  '.',
				{X: 0, Y: 1}:   '.',
				{X: 1, Y: 1}:   '.',
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := applyRules(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestExample1(t *testing.T) {
	actual := findSteadyState(strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`))
	assert.Equal(t, 37, actual)
}

func TestPart1(t *testing.T) {
	actual := findSteadyState(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 2164, actual)
}

func TestLOSOccupiedSeats(t *testing.T) {
	tests := map[string]struct {
		input    string
		loc      common.Point
		expected int
	}{
		"ex1": {
			input: `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`,
			loc:      common.Point{X: 3, Y: 4},
			expected: 8,
		},
		"ex2": {
			input: `.............
.L.L.#.#.#.#.
.............`,
			loc:      common.Point{X: 1, Y: 1},
			expected: 0,
		},
		"ex3": {
			input: `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`,
			loc:      common.Point{X: 3, Y: 3},
			expected: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			layout, max := common.ReadMap(strings.NewReader(test.input))
			actual := countLOSOccupiedSeats(layout, test.loc, max.X, max.Y)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestExample2(t *testing.T) {
	actual := findNewSteadyState(strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`))
	assert.Equal(t, 26, actual)
}

func TestPart2(t *testing.T) {
	actual := findNewSteadyState(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 1974, actual)
}
