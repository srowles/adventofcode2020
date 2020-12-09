package day9

import (
	"io"

	"github.com/srowles/adventofcode2020/common"
)

func findInvalidNumber(reader io.Reader, preamble int) int {
	numbers := common.GetInts(reader)
	for i := preamble; i < len(numbers); i++ {
		ok := findPair(numbers[i-preamble:i], numbers[i])
		if !ok {
			return numbers[i]
		}
	}
	return -1
}

func findPair(numbers []int, val int) bool {
	for i, v := range numbers {
		if v > val {
			continue
		}
		for _, o := range numbers[i:] {
			if v+o == val {
				return true
			}
		}
	}
	return false
}

func findSequence(reader io.Reader, invalid int) int {
	numbers := common.GetInts(reader)
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] > invalid {
			continue
		}
		for j := i; j >= 0; j-- {
			x := sum(numbers[j:i])
			if x > invalid {
				break
			}
			if x == invalid {
				min, max := minmax(numbers[j:i])
				return min + max
			}
		}
	}
	return -1
}

func sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func minmax(numbers []int) (int, int) {
	max := 0
	min := 999999999999
	for _, i := range numbers {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return min, max
}
