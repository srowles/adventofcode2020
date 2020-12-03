package day3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/srowles/adventofcode2020/common"
)

func TestExample(t *testing.T) {
	grid, size := common.ReadMap(common.ReaderFromFile("example.txt"))
	count := countTrees(grid, size, common.Point{X: 3, Y: 1})
	assert.Equal(t, 7, count)
}

func TestWrapping(t *testing.T) {
	grid, size := common.ReadMap(strings.NewReader(
		`.#.......#.
..##.......
.......#..#
...#.......
###..#.#.##
#..##...#..
#..#..#....`))
	count := countTrees(grid, size, common.Point{X: 3, Y: 1})
	assert.Equal(t, 3, count)
}

func TestPart1(t *testing.T) {
	grid, size := common.ReadMap(common.ReaderFromFile("input.txt"))
	count := countTrees(grid, size, common.Point{X: 3, Y: 1})
	assert.Equal(t, 209, count)
}

func TestPart2Example(t *testing.T) {
	grid, size := common.ReadMap(common.ReaderFromFile("example.txt"))
	count1 := countTrees(grid, size, common.Point{X: 1, Y: 1})
	count2 := countTrees(grid, size, common.Point{X: 3, Y: 1})
	count3 := countTrees(grid, size, common.Point{X: 5, Y: 1})
	count4 := countTrees(grid, size, common.Point{X: 7, Y: 1})
	count5 := countTrees(grid, size, common.Point{X: 1, Y: 2})
	mul := count1 * count2 * count3 * count4 * count5
	assert.Equal(t, 336, mul)
}

func TestPart2(t *testing.T) {
	grid, size := common.ReadMap(common.ReaderFromFile("input.txt"))
	count1 := countTrees(grid, size, common.Point{X: 1, Y: 1})
	count2 := countTrees(grid, size, common.Point{X: 3, Y: 1})
	count3 := countTrees(grid, size, common.Point{X: 5, Y: 1})
	count4 := countTrees(grid, size, common.Point{X: 7, Y: 1})
	count5 := countTrees(grid, size, common.Point{X: 1, Y: 2})
	mul := count1 * count2 * count3 * count4 * count5
	assert.Equal(t, 1574890240, mul)
}
