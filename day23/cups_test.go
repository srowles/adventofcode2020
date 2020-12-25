package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	actual := process(`389125467`, 10)
	assert.Equal(t, 92658374, actual)
}

func TestExample2(t *testing.T) {
	actual := process(`389125467`, 100)
	assert.Equal(t, 67384529, actual)
}

func TestPart1(t *testing.T) {
	actual := process(`739862541`, 100)
	assert.Equal(t, 94238657, actual)
}

func TestExample3(t *testing.T) {
	actual := processPart2(`389125467`, 10000000)
	assert.Equal(t, 149245887792, actual)
}

func TestPart2(t *testing.T) {
	actual := processPart2(`739862541`, 10000000)
	assert.Equal(t, 3072905352, actual)
}
