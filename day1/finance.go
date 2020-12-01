package day1

import (
	"io"

	"github.com/srowles/adventofcode2020/common"
)

func part1(reader io.Reader) int {
	input := common.GetInts(reader)
	// find the 2 numbers in input that add up to 2020
	var a, b int
	for _, x := range input {
		a = x
		for _, y := range input {
			if a+y == 2020 {
				b = y
				// then multiply together for result
				return a * b
			}
		}
	}
	return 0
}

func part2(reader io.Reader) int {
	input := common.GetInts(reader)
	// find the 3 numbers in input that add up to 2020
	for _, x := range input {
		for _, y := range input {
			for _, z := range input {
				if x+y+z == 2020 {
					// then multiply together for result
					return x * y * z
				}
			}
		}
	}
	return 0
}
