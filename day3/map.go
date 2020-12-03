package day3

import (
	"github.com/srowles/adventofcode2020/common"
)

func countTrees(grid map[common.Point]rune, size common.Point, slope common.Point) int {
	location := common.Point{}
	count := 0
	for {
		if location.Y >= size.Y {
			break
		}
		location.X += slope.X
		location.Y += slope.Y
		if location.X > size.X {
			location.X = location.X - size.X - 1
		}
		if grid[location] == '#' {
			count++
		}
	}
	return count
}
