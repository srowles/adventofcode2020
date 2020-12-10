package day10

import (
	"io"
	"sort"

	"github.com/srowles/adventofcode2020/common"
)

func process(reader io.Reader) int {
	adapters := common.GetInts(reader)
	sort.Ints(adapters)

	all := append([]int{0}, adapters...)
	all = append(all, all[len(all)-1]+3)
	var ones, threes int
	for i, jolt := range all[:len(all)-1] {
		if all[i+1] == jolt+1 {
			ones++
		}
		if all[i+1] == jolt+3 {
			threes++
		}

	}
	return ones * threes
}

func process2(reader io.Reader) int {
	adapters := common.GetInts(reader)
	sort.Ints(adapters)

	all := append([]int{0}, adapters...)
	all = append(all, all[len(all)-1]+3)

	counts := make(map[int]int)
	counts[0] = 1
	for _, jolt := range all[1:] {
		counts[jolt] = counts[jolt-1] + counts[jolt-2] + counts[jolt-3]
	}
	return counts[all[len(all)-1]]
}
