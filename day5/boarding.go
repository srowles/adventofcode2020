package day5

import (
	"io"

	"github.com/srowles/adventofcode2020/common"
)

func process(instructions string) int {
	answer := 0
	for _, r := range instructions {
		answer = answer << 1
		if r == 'F' || r == 'L' {
			continue
		} else {
			answer++
		}
	}

	return answer
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
