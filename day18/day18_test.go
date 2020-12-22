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

func TestExample21(t *testing.T) {
	actual := process2(strings.NewReader(`1 + 2 * 3 + 4 * 5 + 6`))
	assert.Equal(t, 231, actual)
}

func TestExample22(t *testing.T) {
	actual := process2(strings.NewReader(`1 + (2 * 3) + (4 * (5 + 6))`))
	assert.Equal(t, 51, actual)
}

func TestExample23(t *testing.T) {
	actual := process2(strings.NewReader(`2 * 3 + (4 * 5)`))
	assert.Equal(t, 46, actual)
}

func TestExample24(t *testing.T) {
	actual := process2(strings.NewReader(`5 + (8 * 3 + 9 + 3 * 4 * 3)`))
	assert.Equal(t, 1445, actual)
}

func TestExample25(t *testing.T) {
	actual := process2(strings.NewReader(`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`))
	assert.Equal(t, 669060, actual)
}

func TestExample26(t *testing.T) {
	actual := process2(strings.NewReader(`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`))
	assert.Equal(t, 23340, actual)
}

func TestSample1(t *testing.T) {
	actual := process2(strings.NewReader(`(4 + (4 * 9)) + 8 + 9`))
	assert.Equal(t, 57, actual)
}

func TestPart2(t *testing.T) {
	actual := process2(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 362464596624526, actual)
}
