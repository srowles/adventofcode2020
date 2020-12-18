package common

import (
	"fmt"
	"strconv"
)

// MustInt panic's if the string is not an int
func MustInt(s string) int {
	a, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("%s is not an int", s))
	}
	return a
}
