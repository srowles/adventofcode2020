package day17

import (
	"fmt"
	"io"

	"github.com/srowles/adventofcode2020/common"
)

func simulate(reader io.Reader, cycles int) int {
	grid, _ := common.ReadMap(reader)
	print(grid)
	for i := 0; i < cycles; i++ {
		grid = sim(grid)
	}
	print(grid)
	count := 0
	for _, p := range grid {
		if p == '#' {
			count++
		}
	}
	return count
}

func simulate4(reader io.Reader, cycles int) int {
	grid, _ := common.ReadMap(reader)
	print(grid)
	for i := 0; i < cycles; i++ {
		grid = sim4(grid)
	}
	count := 0
	for _, p := range grid {
		if p == '#' {
			count++
		}
	}
	return count
}

func print(grid map[common.Point]rune) {
	min, max := findSize(grid)
	for z := min.Z; z <= max.Z; z++ {
		for y := min.Y; y <= max.Y; y++ {
			for x := min.X; x <= max.X; x++ {
				r := grid[common.Point{X: x, Y: y, Z: z}]
				if r == 0 {
					r = '.'
				}
				fmt.Print(string(r))
			}
			fmt.Println()
		}
	}
}

func sim(grid map[common.Point]rune) map[common.Point]rune {
	result := make(map[common.Point]rune)
	min, max := findSize(grid)
	min = common.Point{X: min.X - 1, Y: min.Y - 1, Z: min.Z - 1}
	max = common.Point{X: max.X + 1, Y: max.Y + 1, Z: max.Z + 1}
	for z := min.Z; z <= max.Z; z++ {
		for y := min.Y; y <= max.Y; y++ {
			for x := min.X; x <= max.X; x++ {
				p := common.Point{X: x, Y: y, Z: z}
				activeNeighbours := countNeighbours(grid, p)
				if grid[p] == '#' {
					if activeNeighbours == 2 || activeNeighbours == 3 {
						result[p] = '#'
					}
				} else {
					if activeNeighbours == 3 {
						result[p] = '#'
					}
				}
			}
		}
	}

	return result
}

func sim4(grid map[common.Point]rune) map[common.Point]rune {
	result := make(map[common.Point]rune)
	min, max := findSize(grid)
	min = common.Point{X: min.X - 1, Y: min.Y - 1, Z: min.Z - 1, W: min.W - 1}
	max = common.Point{X: max.X + 1, Y: max.Y + 1, Z: max.Z + 1, W: max.W + 1}
	for w := min.W; w <= max.W; w++ {
		for z := min.Z; z <= max.Z; z++ {
			for y := min.Y; y <= max.Y; y++ {
				for x := min.X; x <= max.X; x++ {
					p := common.Point{X: x, Y: y, Z: z, W: w}
					activeNeighbours := countNeighbours(grid, p)
					if grid[p] == '#' {
						if activeNeighbours == 2 || activeNeighbours == 3 {
							result[p] = '#'
						}
					} else {
						if activeNeighbours == 3 {
							result[p] = '#'
						}
					}
				}
			}
		}
	}

	return result
}

var deltas []common.Point
var deltas4 []common.Point

func init() {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						// ignore us
						continue
					}
					deltas = append(deltas, common.Point{X: x, Y: y, Z: z, W: w})
				}
			}
		}
	}
}

func countNeighbours(grid map[common.Point]rune, loc common.Point) int {
	count := 0
	for _, d := range deltas {
		p := common.Point{X: loc.X + d.X, Y: loc.Y + d.Y, Z: loc.Z + d.Z, W: loc.W + d.W}
		if grid[p] == '#' {
			count++
		}
	}

	return count
}

func findSize(grid map[common.Point]rune) (common.Point, common.Point) {
	minx, miny, minz, minw := 9999999999, 9999999999, 9999999999, 9999999999
	var maxx, maxy, maxz, maxw int
	for p := range grid {
		if p.X < minx {
			minx = p.X
		}
		if p.Y < miny {
			miny = p.Y
		}
		if p.Z < minz {
			minz = p.Z
		}
		if p.W < minw {
			minw = p.W
		}
		if p.X > maxx {
			maxx = p.X
		}
		if p.Y > maxy {
			maxy = p.Y
		}
		if p.Z > maxz {
			maxz = p.Z
		}
		if p.W > maxw {
			maxw = p.W
		}
	}
	return common.Point{X: minx, Y: miny, Z: minz, W: minw}, common.Point{X: maxx, Y: maxy, Z: maxz, W: maxw}
}
