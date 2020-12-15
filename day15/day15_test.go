package day15

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	actual := processNumbers(strings.NewReader(`0,3,6`), 2020)
	assert.Equal(t, 436, actual)
}

func TestPart1(t *testing.T) {
	actual := processNumbers(strings.NewReader(`0,13,16,17,1,10,6`), 2020)
	assert.Equal(t, 276, actual)
}

func TestPart2(t *testing.T) {
	actual := processNumbers(strings.NewReader(`0,13,16,17,1,10,6`), 30000000)
	assert.Equal(t, 31916, actual)
}
