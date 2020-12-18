package day18

import (
	"io"
	"unicode"

	"github.com/srowles/adventofcode2020/common"
)

func process(reader io.Reader) int {
	lines := common.StringListFromReader(reader)
	result := 0
	for _, line := range lines {
		result += parseSum(line)
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
		if unicode.IsSpace(part) {
			continue
		}
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
