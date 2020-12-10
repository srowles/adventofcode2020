package day10

import (
	"strings"
	"testing"

	"github.com/srowles/adventofcode2020/common"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	actual := process(strings.NewReader(`16
10
15
5
1
11
7
19
6
12
4`))
	assert.Equal(t, 7*5, actual)
}

func TestExample2(t *testing.T) {
	actual := process(strings.NewReader(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`))
	assert.Equal(t, 22*10, actual)
}

func TestPart1(t *testing.T) {
	actual := process(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 2244, actual)
}

func TestPart2Example1(t *testing.T) {
	actual := process2(strings.NewReader(`16
10
15
5
1
11
7
19
6
12
4`))
	assert.Equal(t, 8, actual)
}

func TestPart2Example2(t *testing.T) {
	actual := process2(strings.NewReader(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`))
	assert.Equal(t, 19208, actual)
}

func TestPart2(t *testing.T) {
	actual := process2(common.ReaderFromFile("input.txt"))
	assert.Equal(t, 3947645370368, actual)
}
