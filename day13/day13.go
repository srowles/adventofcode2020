package day13

import (
	"io"
	"strconv"
	"strings"

	"github.com/srowles/adventofcode2020/common"
)

func processInstructions(reader io.Reader) int {
	list := common.StringListFromReader(reader)
	earliestDepart, _ := strconv.Atoi(list[0])
	busses := strings.Split(list[1], ",")
	shortestWait := 99999999999
	shortestBus := 0
	for _, bus := range busses {
		if bus == "x" {
			continue
		}
		busNumber, _ := strconv.Atoi(bus)
		wait := busNumber - (earliestDepart % busNumber)
		if wait < shortestWait {
			shortestWait = wait
			shortestBus = busNumber
		}
	}

	return shortestBus * shortestWait
}

func processInstructions2(reader io.Reader) int {
	list := common.StringListFromReader(reader)
	busses := strings.Split(list[1], ",")
	timestamp := 0
	step := 1
	for idx, bus := range busses {
		if bus == "x" {
			continue
		}
		busNumber, _ := strconv.Atoi(bus)

		// walk forward by the current step size until we find
		// a time that matches with the new bus we are trying to locate
		for (timestamp+idx)%busNumber != 0 {
			timestamp += step
		}
		// busses are prime, so we can multiply our
		// current bus number by the current step size
		// to find the next lot of steps to move forward in
		step *= busNumber
	}

	return timestamp
}
