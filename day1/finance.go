package day1

import (
	"io"

	"github.com/srowles/adventofcode2020/common"
)

func part1(reader io.Reader) int {
	input := common.GetInts(reader)
	// find the 2 numbers in input that add up to 2020
	for i, x := range input {
		for _, y := range input[i:] {
			if x+y == 2020 {
				// then multiply together for result
				return x * y
			}
		}
	}
	return 0
}

func part2(reader io.Reader) int {
	input := common.GetInts(reader)
	// find the 3 numbers in input that add up to 2020
	for i, x := range input {
		for ii, y := range input[i:] {
			for _, z := range input[ii:] {
				if x+y+z == 2020 {
					// then multiply together for result
					return x * y * z
				}
			}
		}
	}
	return 0
}
