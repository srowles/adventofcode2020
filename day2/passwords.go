package day2

import (
	"fmt"
	"strings"
)

// Rule holds letter and min/max
type Rule struct {
	min    int
	max    int
	letter string
}

func parseRule(rule string) Rule {
	var min, max int
	var letter string
	fmt.Sscanf(rule, "%d-%d %s", &min, &max, &letter)
	return Rule{
		min:    min,
		max:    max,
		letter: letter,
	}
}

func parseLine(line string) (Rule, string) {
	parts := strings.Split(line, ": ")
	return parseRule(parts[0]), parts[1]
}

func checkPassword(rule Rule, password string) bool {
	c := strings.Count(password, rule.letter)
	if c >= rule.min && c <= rule.max {
		return true
	}

	return false
}

func processPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		r, p := parseLine(line)
		if checkPassword(r, p) {
			count++
		}
	}
	return count
}

func checkPassword2(rule Rule, password string) bool {
	first := string(password[rule.min-1])
	second := string(password[rule.max-1])
	if first == rule.letter && second != rule.letter {
		return true

	}
	if second == rule.letter && first != rule.letter {
		return true
	}
	return false
}

func processPasswords2(lines []string) int {
	count := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		r, p := parseLine(line)
		if checkPassword2(r, p) {
			count++
		}
	}
	return count
}
