package day5

import (
	"testing"

	"github.com/srowles/adventofcode2020/common"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {

	tests := map[string]struct {
		instructions string
		answer       int
	}{
		"ex1": {
			instructions: "FBFBBFFRLR",
			answer:       357,
		},
		"ex2": {
			instructions: "BFFFBBFRRR",
			answer:       567,
		},
		"ex3": {
			instructions: "FFFBBBFRRR",
			answer:       119,
		},
		"ex4": {
			instructions: "BBFFBBFRLL",
			answer:       820,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ticketID := process(test.instructions)
			assert.Equal(t, test.answer, ticketID)
		})
	}
}

func TestPart1(t *testing.T) {
	largestSeatID := processTickets(common.ReaderFromFile(`input.txt`))
	assert.Equal(t, 965, largestSeatID)
}

func TestPart2(t *testing.T) {
	seatID := findTicket(common.ReaderFromFile(`input.txt`))
	assert.Equal(t, 524, seatID)
}
