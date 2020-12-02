package day2

import (
	"testing"

	"github.com/srowles/adventofcode2020/common"

	"github.com/stretchr/testify/assert"
)

func TestParseRule(t *testing.T) {
	tests := map[string]struct {
		input string
		rule  Rule
	}{
		"1": {
			input: "1-3 a",
			rule: Rule{
				letter: "a",
				min:    1,
				max:    3,
			},
		},
		"2": {
			input: "12-34 z",
			rule: Rule{
				letter: "z",
				min:    12,
				max:    34,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := parseRule(test.input)
			assert.Equal(t, test.rule, r)
		})
	}
}

func TestParseLine(t *testing.T) {
	tests := map[string]struct {
		input string
		rule  Rule
		pass  string
	}{
		"1": {
			input: "2-6 f: abcdef",
			rule: Rule{
				letter: "f",
				min:    2,
				max:    6,
			},
			pass: "abcdef",
		},
		"2": {
			input: "2-126 y: poiuyt",
			rule: Rule{
				letter: "y",
				min:    2,
				max:    126,
			},
			pass: "poiuyt",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r, p := parseLine(test.input)
			assert.Equal(t, test.rule, r)
			assert.Equal(t, test.pass, p)
		})
	}
}

func TestPasswords(t *testing.T) {
	count := processPasswords(common.StringListFromFile(common.ReaderFromFile("input1.txt")))
	assert.Equal(t, 560, count)
}

func TestCheckPassword2(t *testing.T) {
	tests := map[string]struct {
		input string
		valid bool
	}{
		"1": {
			input: "2-6 f: abcdez",
			valid: false,
		},
		"2": {
			input: "2-3 y: pyoiut",
			valid: true,
		},
		"3": {
			input: "2-6 z: pyoiuzhj",
			valid: true,
		},
		"4": {
			input: "2-6 z: pzoiuzhj",
			valid: false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r, p := parseLine(test.input)
			ok := checkPassword2(r, p)
			assert.Equal(t, test.valid, ok)
		})
	}
}

func TestPasswords2(t *testing.T) {
	count := processPasswords2(common.StringListFromFile(common.ReaderFromFile("input1.txt")))
	assert.Equal(t, 303, count)
}
