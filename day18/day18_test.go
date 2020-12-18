package day18

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	actual := process(strings.NewReader(`1 + 2 * 3 + 4 * 5 + 6`))
	assert.Equal(t, 71, actual)
}

func TestExample2(t *testing.T) {
	actual := process(strings.NewReader(`1 + (2 * 3) + (4 * (5 + 6))`))
	assert.Equal(t, 51, actual)
}

func TestExample3(t *testing.T) {
	actual := process(strings.NewReader(`2 * 3 + (4 * 5)`))
	assert.Equal(t, 26, actual)
}

func TestExample4(t *testing.T) {
	actual := process(strings.NewReader(`5 + (8 * 3 + 9 + 3 * 4 * 3)`))
	assert.Equal(t, 437, actual)
}

func TestExample5(t *testing.T) {
	actual := process(strings.NewReader(`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`))
	assert.Equal(t, 12240, actual)
}

func TestExample6(t *testing.T) {
	actual := process(strings.NewReader(`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`))
	assert.Equal(t, 13632, actual)
}

func TestPart1(t *testing.T) {
	actual := process(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 69490582260, actual)
}
