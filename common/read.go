package common

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// StringListFromReader parses new line separated text into a list of strings
func StringListFromReader(reader io.Reader) []string {
	data := mustReadStringData(reader)
	lines := strings.Split(data, "\n")
	return lines
}

// ReaderFromFile creates an io.Reader from supplied file or panics
func ReaderFromFile(path string) io.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

// ReaderFromInts will create an io.Reader from the int slice supplied
func ReaderFromInts(input []int) io.Reader {
	var builder strings.Builder
	for _, v := range input {
		builder.WriteString(fmt.Sprintf("%d\n", v))
	}
	return bytes.NewReader([]byte(builder.String()))
}

// IntList comma separated single line list of numbers
func IntList(reader io.Reader) []int {
	data := mustReadStringData(reader)
	numbers := strings.Split(data, ",")
	result := make([]int, 0, len(numbers))
	for _, l := range numbers {
		if strings.TrimSpace(l) == "" {
			continue
		}
		result = append(result, MustInt(l))
	}
	return result
}

// GetInts returns a slice of ints from a file
func GetInts(reader io.Reader) []int {
	data := mustReadStringData(reader)
	lines := strings.Split(data, "\n")
	result := make([]int, 0, len(lines))
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		result = append(result, MustInt(l))
	}

	return result
}

// Point is a grid coordinate, X horizontal y vertical
type Point struct {
	X, Y, Z, W int
}

// ReadMap reads data from a read into a sparce map backed grid representation
func ReadMap(reader io.Reader) (map[Point]rune, Point) {
	data := mustReadStringData(reader)
	lines := strings.Split(data, "\n")
	grid := make(map[Point]rune)
	var maxx, maxy int
	for y, l := range lines {
		if l == "" {
			break
		}
		for x, c := range l {
			p := Point{X: x, Y: y}
			if c != '.' {
				grid[p] = c
			}
			maxx = x
			maxy = y
		}
	}

	return grid, Point{maxx, maxy, 0, 0}
}

func mustReadStringData(reader io.Reader) string {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return string(data)
}
