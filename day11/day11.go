package day11

import (
	"fmt"
	"io"

	"github.com/srowles/adventofcode2020/common"
)

func findSteadyState(reader io.Reader) int {
	layout, _ := common.ReadMap(reader)
	newLayout := applyRules(layout)
	for {
		newLayout = applyRules(layout)
		if same(newLayout, layout) {
			break
		}
		for k, v := range newLayout {
			layout[k] = v
		}
	}

	count := 0
	for _, v := range newLayout {
		if v == '#' {
			count++
		}
	}
	return count
}

func findNewSteadyState(reader io.Reader) int {
	layout, max := common.ReadMap(reader)
	newLayout := applyRules2(layout, max.X, max.Y)
	for {
		newLayout = applyRules2(layout, max.X, max.Y)
		if same(newLayout, layout) {
			break
		}
		for k, v := range newLayout {
			layout[k] = v
		}
	}

	count := 0
	for _, v := range newLayout {
		if v == '#' {
			count++
		}
	}
	return count
}

func print(layout map[common.Point]rune) {
	var maxx, maxy int
	for k := range layout {
		if k.X > maxx {
			maxx = k.X
		}
		if k.Y > maxy {
			maxy = k.Y
		}
	}
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			r := layout[common.Point{X: x, Y: y}]
			if r == 0 {
				r = '.'
			}
			fmt.Print(string(r))
		}
		fmt.Println()
	}
	fmt.Println()
}

func applyRules(layout map[common.Point]rune) map[common.Point]rune {
	result := make(map[common.Point]rune, len(layout))
	for k, v := range layout {
		// by default nothing happens
		result[k] = v

		// apply rules based on current seat value
		switch v {
		case 'L':
			adjacent := countAdjacentOccupiedSeats(layout, k)
			if adjacent == 0 {
				result[k] = '#'
			}
		case '#':
			adjacent := countAdjacentOccupiedSeats(layout, k)
			if adjacent >= 4 {
				result[k] = 'L'
			}
		}
	}
	return result
}

func applyRules2(layout map[common.Point]rune, maxx, maxy int) map[common.Point]rune {
	result := make(map[common.Point]rune, len(layout))
	for k, v := range layout {
		// by default nothing happens
		result[k] = v

		// apply rules based on current seat value
		switch v {
		case 'L':
			adjacent := countLOSOccupiedSeats(layout, k, maxx, maxy)
			if adjacent == 0 {
				result[k] = '#'
			}
		case '#':
			adjacent := countLOSOccupiedSeats(layout, k, maxx, maxy)
			if adjacent >= 5 {
				result[k] = 'L'
			}
		}
	}
	return result
}

var adjacent = []common.Point{
	{X: -1, Y: -1},
	{X: 0, Y: -1},
	{X: 1, Y: -1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: -1, Y: 1},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
}

func countAdjacentOccupiedSeats(layout map[common.Point]rune, loc common.Point) int {
	count := 0
	for _, mod := range adjacent {
		p := common.Point{X: loc.X + mod.X, Y: loc.Y + mod.Y}
		if layout[p] == '#' {
			count++
		}
	}
	return count
}
func countLOSOccupiedSeats(layout map[common.Point]rune, loc common.Point, maxx, maxy int) int {
	count := 0
	for _, mod := range adjacent {
		p := common.Point{X: loc.X, Y: loc.Y}
		for {
			p = common.Point{X: p.X + mod.X, Y: p.Y + mod.Y}
			if p.X < 0 || p.Y < 0 || p.X > maxx || p.Y > maxy {
				break
			}
			if layout[p] == '#' {
				count++
				break
			}
			if layout[p] == 'L' {
				break
			}
		}
	}
	return count
}

func same(current, new map[common.Point]rune) bool {
	for k, v := range current {
		if new[k] != v {
			return false
		}
	}
	return true
}
