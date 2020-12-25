package day23

import (
	"container/ring"
	"fmt"
	"strings"

	"github.com/srowles/adventofcode2020/common"
)

func process(order string, iterations int) int {
	cups := ring.New(len(order))
	lookup := make(map[int]*ring.Ring, len(order))
	root := cups
	maxCup := 0
	minCup := 99999999
	for _, n := range order {
		cup := common.MustInt(string(n))
		if cup > maxCup {
			maxCup = cup
		}
		if cup < minCup {
			minCup = cup
		}
		cups.Value = cup
		lookup[cup] = cups
		cups = cups.Next()
	}
	currentCup := root
	for i := 0; i < iterations; i++ {
		removed := currentCup.Unlink(3)
		nextCupNumber := currentCup.Value.(int) - 1
		if nextCupNumber < minCup {
			nextCupNumber = maxCup
		}
		for inRemoved(nextCupNumber, removed) {
			nextCupNumber--
			if nextCupNumber < minCup {
				nextCupNumber = maxCup
			}
		}

		insert := lookup[nextCupNumber]
		insert.Link(removed)
		currentCup = currentCup.Next()
	}

	start := lookup[1].Unlink(len(order) - 1)
	var result strings.Builder
	for j := 0; j < start.Len(); j++ {
		result.WriteString(fmt.Sprintf("%d", start.Value.(int)))
		start = start.Next()
	}
	return common.MustInt(result.String())
}

func processPart2(order string, iterations int) int {
	cups := ring.New(1000000)
	root := cups
	maxCup := 0
	minCup := 99999999
	lookup := make(map[int]*ring.Ring, 1000000)
	for _, n := range order {
		cup := common.MustInt(string(n))
		if cup > maxCup {
			maxCup = cup
		}
		if cup < minCup {
			minCup = cup
		}
		cups.Value = cup
		lookup[cup] = cups
		cups = cups.Next()
	}
	for i := maxCup + 1; i <= 1000000; i++ {
		cups.Value = i
		lookup[i] = cups
		cups = cups.Next()
	}
	maxCup = 1000000
	currentCup := root
	for i := 0; i < iterations; i++ {
		removed := currentCup.Unlink(3)
		nextCupNumber := currentCup.Value.(int) - 1
		if nextCupNumber < minCup {
			nextCupNumber = maxCup
		}
		for inRemoved(nextCupNumber, removed) {
			nextCupNumber--
			if nextCupNumber < minCup {
				nextCupNumber = maxCup
			}
		}

		insert := lookup[nextCupNumber]
		insert.Link(removed)
		currentCup = currentCup.Next()
	}
	one := lookup[1]
	part1 := one.Next().Value.(int)
	part2 := one.Next().Next().Value.(int)
	return part1 * part2
}

func inRemoved(num int, r *ring.Ring) bool {
	found := false
	r.Do(func(n interface{}) {
		if n == num {
			found = true
		}
	})
	return found
}

func print(root *ring.Ring) {
	var ints []int
	for j := 0; j < root.Len(); j++ {
		ints = append(ints, root.Value.(int))
		root = root.Next()
	}
	fmt.Println(ints)
}
