package day13

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	actual := processInstructions(strings.NewReader(`939
7,13,x,x,59,x,31,19`))
	assert.Equal(t, 295, actual)
}

func TestPart1(t *testing.T) {
	actual := processInstructions(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 3865, actual)
}

func TestExample2(t *testing.T) {
	actual := processInstructions2(strings.NewReader(`939
7,13,x,x,59,x,31,19`))
	assert.Equal(t, 1068781, actual)
}

func TestExample3(t *testing.T) {
	actual := processInstructions2(strings.NewReader(`939
17,x,13,19`))
	assert.Equal(t, 3417, actual)
}

func TestExample4(t *testing.T) {
	actual := processInstructions2(strings.NewReader(`939
67,7,59,61`))
	assert.Equal(t, 754018, actual)
}

func TestExample5(t *testing.T) {
	actual := processInstructions2(strings.NewReader(`939
67,x,7,59,61`))
	assert.Equal(t, 779210, actual)
}

func TestExample6(t *testing.T) {
	actual := processInstructions2(strings.NewReader(`939
1789,37,47,1889`))
	assert.Equal(t, 1202161486, actual)
}

func TestPart2(t *testing.T) {
	actual := processInstructions2(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 415579909629976, actual)
}
