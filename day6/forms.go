package day6

import (
	"io"

	"github.com/srowles/adventofcode2020/common"
)

type form struct {
	people  int
	answers map[string]int
}

func process(reader io.Reader) (int, int) {
	var forms []form
	lines := common.StringListFromFile(reader)
	p := form{answers: make(map[string]int)}
	for _, l := range lines {
		if l == "" {
			// finish parsing a form, append it
			forms = append(forms, p)
			p = form{answers: make(map[string]int)}
			continue
		}
		p.people++
		for _, r := range l {
			p.answers[string(r)]++
		}
	}

	count := 0
	allCount := 0
	for _, f := range forms {
		for _, v := range f.answers {
			if v == f.people {
				allCount++
			}
		}
		count += len(f.answers)
	}
	return count, allCount
}
