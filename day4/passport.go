package day4

import (
	"io"
	"strconv"
	"strings"

	"github.com/srowles/adventofcode2020/common"
)

// Passport is a type alias for map of field to value
type Passport map[string]string

// Valid checks if a passport is valid, with optional detailed check
func (p Passport) Valid(detailed bool) bool {
	count := 0
	for key := range p {
		switch key {
		case "byr":
			count++
		case "iyr":
			count++
		case "eyr":
			count++
		case "hgt":
			count++
		case "hcl":
			count++
		case "ecl":
			count++
		case "pid":
			count++
		}
	}
	if count != 7 {
		return false
	}
	// cid missing is OK
	if detailed {
		return p.detailed()
	}

	return true
}

// validate the fields in detail
func (p Passport) detailed() bool {
	for key, val := range p {
		switch key {
		case "byr":
			if !validByr(val) {
				return false
			}
		case "iyr":
			if !validIyr(val) {
				return false
			}
		case "eyr":
			if !validEyr(val) {
				return false
			}
		case "hgt":
			if !validHgt(val) {
				return false
			}
		case "hcl":
			if !validHcl(val) {
				return false
			}
		case "ecl":
			if !validEcl(val) {
				return false
			}
		case "pid":
			if !validPid(val) {
				return false
			}
		}
	}

	return true
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func validByr(val string) bool {
	if len(val) != 4 {
		return false
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if num < 1920 || num > 2002 {
		return false
	}
	return true
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func validIyr(val string) bool {
	if len(val) != 4 {
		return false
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if num < 2010 || num > 2020 {
		return false
	}
	return true
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func validEyr(val string) bool {
	if len(val) != 4 {
		return false
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if num < 2020 || num > 2030 {
		return false
	}
	return true
}

// hgt (Height) - a number followed by either cm or in:
//    If cm, the number must be at least 150 and at most 193.
//    If in, the number must be at least 59 and at most 76.
func validHgt(val string) bool {
	if len(val) < 2 {
		return false
	}

	typ := val[len(val)-2:]
	if typ == "cm" {
		h, err := strconv.Atoi(val[0 : len(val)-2])
		if err != nil {
			return false
		}
		if h < 150 || h > 193 {
			return false
		}
		return true
	}

	if typ == "in" {
		h, err := strconv.Atoi(val[0 : len(val)-2])
		if err != nil {
			return false
		}
		if h < 59 || h > 76 {
			return false
		}
		return true
	}

	return false
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func validHcl(val string) bool {
	if len(val) != 7 {
		return false
	}

	val = val[1:]
	for _, c := range val {
		if c < 48 || (c > 57 && c < 97) || c > 102 {
			return false
		}
	}

	return true
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func validEcl(val string) bool {
	switch val {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	}

	return false
}

// (Passport ID) - a nine-digit number, including leading zeroes
func validPid(val string) bool {
	if len(val) != 9 {
		return false
	}

	for _, c := range val {
		if c < 48 || c > 57 {
			return false
		}
	}

	return true
}

// Passports is a list of passports to be checking
type Passports struct {
	passports []Passport
}

func parse(reader io.Reader) *Passports {
	var passports []Passport
	lines := common.StringListFromFile(reader)
	p := make(map[string]string)
	for _, l := range lines {
		if l == "" {
			// finish parsing a passport, append it
			passports = append(passports, p)
			p = make(map[string]string)
			continue
		}
		fields := strings.Fields(l)
		for _, f := range fields {
			parts := strings.Split(f, ":")
			p[parts[0]] = parts[1]
		}
	}

	return &Passports{
		passports: passports,
	}
}

// ValidCount returns the number of valid passports, with optional detailed check
func (p *Passports) ValidCount(detailed bool) int {
	count := 0
	for _, passport := range p.passports {
		if passport.Valid(detailed) {
			count++
		}
	}
	return count
}
