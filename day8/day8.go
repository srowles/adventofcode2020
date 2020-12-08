package day8

import (
	"fmt"
	"io"

	"github.com/srowles/adventofcode2020/common"
)

type computer struct {
	insructionPointer int
	accumulator       int
}

type instruction struct {
	instr string
	val   int
}

func (c *computer) Process(instructions []instruction) (bool, int) {
	count := make(map[int]int)
	c.insructionPointer = 0
	for {
		if c.insructionPointer >= len(instructions) {
			return false, c.accumulator
		}
		if count[c.insructionPointer] == 1 {
			return true, c.accumulator
		}
		count[c.insructionPointer]++
		switch instructions[c.insructionPointer].instr {
		case "nop":
		case "acc":
			c.accumulator += instructions[c.insructionPointer].val
		case "jmp":
			c.insructionPointer += instructions[c.insructionPointer].val
			continue
		}
		c.insructionPointer++
	}
}

func getInstructions(reader io.Reader) []instruction {
	var instructions []instruction
	for _, line := range common.StringListFromReader(reader) {
		if line == "" {
			continue
		}
		var instr string
		var val int
		fmt.Sscanf(line, "%s %d", &instr, &val)
		instructions = append(instructions, instruction{instr: instr, val: val})
	}
	return instructions
}

func process(reader io.Reader) int {
	comp := &computer{}
	_, val := comp.Process(getInstructions(reader))

	return val
}

func process2(reader io.Reader) int {
	instructions := getInstructions(reader)
	modified := make([]instruction, len(instructions))
	for i := 0; i < len(instructions); i++ {
		copy(modified, instructions)
		switch modified[i].instr {
		case "jmp":
			modified[i] = instruction{
				instr: "nop",
				val:   modified[i].val,
			}
		case "nop":
			modified[i] = instruction{
				instr: "jmp",
				val:   modified[i].val,
			}
		}

		comp := &computer{}
		looped, val := comp.Process(modified)
		if !looped {
			return val
		}
	}

	return -1
}
