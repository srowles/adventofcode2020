package day18

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/srowles/adventofcode2020/common"
)

func process(reader io.Reader) int {
	lines := common.StringListFromReader(reader)
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		result += parseSum(strings.ReplaceAll(line, " ", ""))
	}

	return result
}

func process2(reader io.Reader) int {
	lines := common.StringListFromReader(reader)
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		val := parseSumWithPrecedence(line)
		result += val
	}

	return result
}

type operator int

const (
	add operator = iota
	mul
)

func parseSum(sum string) int {
	opStack := new(common.IntStack)
	opStack.Push(int(add))
	stack := new(common.IntStack)
	stack.Push(0)
	for _, part := range sum {
		v := int(part - '0')
		switch part {
		case '(':
			stack.Push(0)
			opStack.Push(int(add))
		case ')':
			val := stack.Pop()
			v := stack.Pop()
			// number
			switch operator(opStack.Pop()) {
			case add:
				val += v
			case mul:
				val *= v
			}
			stack.Push(val)
		case '+':
			opStack.Push(int(add))
		case '*':
			opStack.Push(int(mul))
		default:
			val := stack.Pop()
			// number
			switch operator(opStack.Pop()) {
			case add:
				val += v
			case mul:
				val *= v
			}
			stack.Push(val)
		}
	}

	val := stack.Pop()
	return val
}

func parseSumWithPrecedence(sum string) int {
	line := []rune(sum)
	// scan backwards through the sum finding bracket groups
	// to recursively process
	end := 0

	for i := len(line) - 1; i >= 0; i-- {
		switch line[i] {
		case ' ':
			continue
		case ')':
			end = i
		case '(':
			v := parseSumWithPrecedence(string(line[i+1 : end]))
			val := []rune(fmt.Sprintf("%d", v))
			tail := append(val, line[end+1:]...)
			line = append(line[:i], tail...)
			for p := i; p < len(line); p++ {
				if line[p] == ')' {
					end = p
				}
			}
		}
	}
	sum = string(line)
	// next do addition first scanning forwards
outer:
	for strings.Contains(sum, "+") {
		prev := 0
		var op string
		for _, f := range strings.Fields(sum) {
			n, _ := strconv.Atoi(f)
			switch f {
			case "+":
				op = "+"
			case "*":
				op = "*"
			default:
				if op == "+" {
					v := prev + n
					sum = strings.Replace(sum, fmt.Sprintf("%d + %d", prev, n), fmt.Sprintf("%d", v), 1)
					continue outer
				}
				prev = n
			}
		}
	}

	// perform last multiplications, nothing else left
	result := 1
	for _, f := range strings.Fields(sum) {
		n, _ := strconv.Atoi(f)
		switch f {
		case "*":
		default:
			result *= n
		}
	}

	return result
}
