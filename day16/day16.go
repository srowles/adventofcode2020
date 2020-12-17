package day16

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/srowles/adventofcode2020/common"
)

func identifyFields(reader io.Reader, prefix string) int {
	myTicket, validTickets, fields, _ := parse(reader)
	allFields := make(map[string]bool)
	for _, f := range fields {
		for k, v := range f {
			allFields[k] = v
		}
	}

	// for my ticket, find the possible fields it can contain using the numbers
	possibleByIndex := make(map[int]map[string]bool)
	for i := 0; i < len(myTicket); i++ {
		fields := make(map[string]bool)
		for k, v := range allFields {
			fields[k] = v
		}
		possibleByIndex[i] = fields
	}
	// using valid tickets, remove possible indexes when they cannot work
	for _, ticket := range validTickets {
		for i, n := range ticket {
			possibleFields := fields[n]
			impossibleFields := make(map[string]bool)
			for f := range possibleByIndex[i] {
				if !possibleFields[f] {
					impossibleFields[f] = true
				}
			}
			for k := range impossibleFields {
				delete(possibleByIndex[i], k)
			}
		}
	}

	// loop tickets finding single fields then removing
	// from the valid set
	assigned := make(map[string]int)
	index := 0
	for len(assigned) < len(myTicket) {
		if len(possibleByIndex[index]) == 1 {
			key := getOnlyKey(possibleByIndex[index])
			assigned[key] = index
			// now remove this one from all other possibles
			for k := range possibleByIndex {
				delete(possibleByIndex[k], key)
			}
		}
		index++
		if index == len(myTicket) {
			index = 0
		}
	}

	fmt.Println(assigned)
	result := 1
	for k, v := range assigned {
		if strings.HasPrefix(k, prefix) {
			result *= myTicket[v]
		}
	}

	return result
}

func getOnlyKey(m map[string]bool) string {
	for k := range m {
		return k
	}
	panic("nokey")
}

func findErrorRate(reader io.Reader) int {
	_, _, _, rate := parse(reader)

	return rate
}

func parse(reader io.Reader) ([]int, [][]int, map[int]map[string]bool, int) {
	lines := common.StringListFromReader(reader)
	validNumbers := make(map[int]bool)
	fields := make(map[int]map[string]bool)
	index := 0
	for i, line := range lines {
		if line == "" {
			index = i
			break
		}
		field, valid := parseRule(line)
		for _, v := range valid {
			validNumbers[v] = true
			if fields[v] == nil {
				fields[v] = make(map[string]bool)
			}
			fields[v][field] = true
		}
	}

	var myTicket []int
	for i := index; i < len(lines); i++ {
		if lines[i] == "your ticket:" {
			index = i + 2
			numbers := strings.Split(lines[i+1], ",")
			for _, num := range numbers {
				n, _ := strconv.Atoi(num)
				myTicket = append(myTicket, n)
			}
			break
		}
	}

	for i := index; i < len(lines); i++ {
		if lines[i] == "nearby tickets:" {
			index = i + 1
			break
		}
	}

	var validTickets [][]int
	rate := 0
	for i := index; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		numbers := strings.Split(lines[i], ",")
		valid := true
		var ticket []int
		for _, num := range numbers {
			n, _ := strconv.Atoi(num)
			ticket = append(ticket, n)
			if !validNumbers[n] {
				valid = false
				rate += n
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	return myTicket, validTickets, fields, rate
}

// parseRule takes a line of text and figures out the rule name and valid values
func parseRule(rule string) (string, []int) {
	parts := strings.Split(rule, ": ")
	sectionName := parts[0]
	ranges := strings.Split(parts[1], " or ")
	var values []int
	for _, r := range ranges {
		var min, max int
		fmt.Sscanf(r, "%d-%d", &min, &max)
		for i := min; i <= max; i++ {
			values = append(values, i)
		}
	}
	return sectionName, values
}
