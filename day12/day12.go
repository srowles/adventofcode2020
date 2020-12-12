package day12

import (
	"io"
	"strconv"

	"github.com/srowles/adventofcode2020/common"
)

type direction int

// east = 0
// south = 1
// west = 2
// north = 3
const (
	east direction = iota
	south
	west
	north
)

func processInstructions(reader io.Reader) int {
	lines := common.StringListFromReader(reader)
	var x, y int
	var dir direction
	for _, line := range lines {
		if line == "" {
			continue
		}
		instruction := line[0:1]
		val, _ := strconv.Atoi(line[1:])
		dir = processDirChange(instruction, val, dir)
		dx, dy := move(instruction, val, dir)
		x += dx
		y += dy
	}
	return abs(x) + abs(y)
}

func processInstructions2(reader io.Reader) int {
	lines := common.StringListFromReader(reader)
	var x, y int
	wx, wy := 10, 1
	for _, line := range lines {
		if line == "" {
			continue
		}
		instruction := line[0:1]
		val, _ := strconv.Atoi(line[1:])
		wx, wy = rotateWaypoint(wx, wy, instruction, val)
		if instruction != "F" {
			dx, dy := move(instruction, val, north)
			wx += dx
			wy += dy
		} else {
			dx, dy := moveShip(val, wx, wy)
			x += dx
			y += dy
		}
	}
	return abs(x) + abs(y)
}

func abs(i int) int {
	if i < 0 {
		return 0 - i
	}
	return i
}

func processDirChange(instruction string, val int, dir direction) direction {
	if instruction == "L" {
		turns := val / 90
		newDir := int(dir) - turns
		if newDir < 0 {
			newDir += 4
		}
		return direction(newDir)
	}
	if instruction == "R" {
		turns := val / 90
		newDir := int(dir) + turns
		if newDir >= 4 {
			newDir -= 4
		}
		return direction(newDir)
	}
	return dir
}

func move(instruction string, val int, dir direction) (int, int) {
	switch instruction {
	case "F":
		switch dir {
		case east:
			return val, 0
		case south:
			return 0, 0 - val
		case west:
			return 0 - val, 0
		case north:
			return 0, val
		}
	case "E":
		return val, 0
	case "S":
		return 0, 0 - val
	case "W":
		return 0 - val, 0
	case "N":
		return 0, val
	}

	return 0, 0
}

func moveShip(val int, x, y int) (int, int) {
	return x * val, y * val
}

func rotateWaypoint(x, y int, instruction string, val int) (int, int) {
	if instruction == "L" {
		turns := val / 90
		for i := 0; i < turns; i++ {
			x, y = y, x
			x = -x
		}
	}
	if instruction == "R" {
		turns := val / 90
		for i := 0; i < turns; i++ {
			x, y = y, x
			y = -y
		}
	}
	return x, y
}
