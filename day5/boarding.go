package day5

import (
	"io"
	"strings"

	"github.com/srowles/adventofcode2020/common"
)

// binary search through the rows with instructions
// Start by considering the whole range, rows 0 through 127.
// F means to take the lower half, keeping rows 0 through 63.
// B means to take the upper half, keeping rows 32 through 63.
// F means to take the lower half, keeping rows 32 through 47.
// B means to take the upper half, keeping rows 40 through 47.
// B keeps rows 44 through 47.
// F keeps rows 44 through 45.
// The final F keeps the lower of the two, row 44.
func split(instructions string, left, right rune, size int) int {
	min := 0
	max := size - 1
	for _, i := range instructions {
		if i == left {
			max -= ((max - min) / 2) + 1
		}
		if i == right {
			min = min + ((max - min) / 2) + 1
		}
	}
	return min
}

func process(instructions string) int {
	var fb, lr strings.Builder
	for _, r := range instructions {
		if r == 'F' || r == 'B' {
			fb.WriteRune(r)
		} else {
			lr.WriteRune(r)
		}
	}

	row := split(fb.String(), 'F', 'B', 128)
	col := split(lr.String(), 'L', 'R', 8)

	return (row * 8) + col
}

func processTickets(reader io.Reader) int {
	rows := common.StringListFromFile(reader)
	maxID := 0
	for _, row := range rows {
		if row == "" {
			continue
		}
		id := process(row)
		if id > maxID {
			maxID = id
		}
	}

	return maxID
}

func findTicket(reader io.Reader) int {
	rows := common.StringListFromFile(reader)
	ids := make(map[int]bool)
	max := 0
	min := 999999999999
	for _, row := range rows {
		if row == "" {
			continue
		}
		id := process(row)
		if id < min {
			min = id
		}
		if id > max {
			max = id
		}
		ids[id] = true
	}

	for i := min; i <= max; i++ {
		if !ids[i] {
			return i
		}
	}

	return -1
}
