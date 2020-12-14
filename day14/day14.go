package day14

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/srowles/adventofcode2020/common"
)

func processInstructions(reader io.Reader) int {
	list := common.StringListFromReader(reader)
	memory := make(map[int]int)
	var andMask, orMask int
	for _, line := range list {
		if line == "" {
			continue
		}
		if strings.Contains(line, "mask") {
			parts := strings.Split(line, " = ")
			andMask, orMask = masks(parts[1])
			continue
		}
		var mem, value int
		fmt.Sscanf(line, "mem[%d] = %d", &mem, &value)
		memory[mem] = apply(value, andMask, orMask)
	}

	return sum(memory)
}

func processInstructions2(reader io.Reader) int {
	list := common.StringListFromReader(reader)
	memory := make(map[int]int)
	var mask string
	for _, line := range list {
		if line == "" {
			continue
		}
		if strings.Contains(line, "mask") {
			parts := strings.Split(line, " = ")
			mask = parts[1]
			continue
		}
		var mem, value int
		fmt.Sscanf(line, "mem[%d] = %d", &mem, &value)
		for _, addr := range expandAddresses(mem, mask) {
			memory[addr] = value
		}
	}

	return sum(memory)
}

// return a list of all possible addresses based
// on mask
// address: 000000000000000000000000000000101010  (decimal 42)
// mask:    000000000000000000000000000000X1001X
// result:  000000000000000000000000000000X1101X
// After applying the mask, four bits are overwritten, three of which are different, and two of which are floating. Floating bits take on every possible combination of values; with two floating bits, four actual memory addresses are written:

// 000000000000000000000000000000011010  (decimal 26)
// 000000000000000000000000000000011011  (decimal 27)
// 000000000000000000000000000000111010  (decimal 58)
// 000000000000000000000000000000111011  (decimal 59)
func expandAddresses(address int, mask string) []int {
	newMask := make([]rune, len(mask))
	for i, r := range fmt.Sprintf("%036b", address) {
		switch mask[i] {
		case '1':
			newMask[i] = '1'
		case '0':
			newMask[i] = r
		case 'X':
			newMask[i] = 'X'
		}
	}
	// iterate varitions
	count := strings.Count(string(newMask), "X")
	combinations := int(math.Exp2(float64(count)))
	result := []int{}
	for i := 0; i < combinations; i++ {
		// combination in binary is the values of X we want replacing
		binary := fmt.Sprintf("%0*b", count, i)
		iterationMask := string(newMask)
		for _, v := range binary {
			iterationMask = strings.Replace(iterationMask, "X", fmt.Sprintf("%c", v), 1)
		}
		val, _ := strconv.ParseInt(iterationMask, 2, 64)
		result = append(result, int(val))
	}

	return result
}

func apply(value, andMask, orMask int) int {
	value &= andMask
	value |= orMask
	return value
}

// value:  000000000000000000000000000000001011  (decimal 11)
// mask:   XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
// result: 000000000000000000000000000001001001  (decimal 73)
func masks(mask string) (int, int) {
	var andMask, orMask int
	for i, r := range mask {
		switch r {
		case 'X':
			andMask |= 1
			orMask |= 0
		case '1':
			andMask |= 1
			orMask |= 1
		case '0':
			andMask |= 0
			orMask |= 0
		}
		if i == len(mask)-1 {
			break
		}

		andMask = andMask << 1
		orMask = orMask << 1
	}

	return andMask, orMask
}

func sum(memory map[int]int) int {
	var sum int
	for _, v := range memory {
		sum += v
	}
	return sum
}
