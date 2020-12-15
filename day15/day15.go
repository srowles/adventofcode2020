package day15

import (
	"io"

	"github.com/srowles/adventofcode2020/common"
)

func processNumbers(reader io.Reader, turns int) int {
	numbers := common.IntList(reader)
	lastNumber := numbers[len(numbers)-1]
	lastTurn := make(map[int][]int)
	count := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		lastTurn[numbers[i]] = append(lastTurn[numbers[i]], i+1)
		count[numbers[i]]++
	}
	for i := len(numbers) + 1; i <= turns; i++ {
		if count[lastNumber] == 1 {
			// first time this number has been
			lastNumber = 0
			lastTurn[0] = append(lastTurn[0], i)
			count[0]++
			continue
		}

		lastNumber = lastTurn[lastNumber][len(lastTurn[lastNumber])-1] - lastTurn[lastNumber][len(lastTurn[lastNumber])-2]
		lastTurn[lastNumber] = append(lastTurn[lastNumber], i)
		count[lastNumber]++
	}
	return lastNumber
}
