package common

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

// GetInts returns a slice of ints from a file
func GetInts(reader io.Reader) []int {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	result := make([]int, 0, len(lines))
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}

	return result
}
