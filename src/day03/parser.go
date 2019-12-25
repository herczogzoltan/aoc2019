package day03

import (
	"bytes"
)

func ParseInput(in []byte) map[int]string {
	results := make(map[int]string, 0)
	lines := bytes.Split(in, []byte("\n"))

	for i := 0; i < len(lines); i++ {
		v := lines[i]
		results[i] = string(v)
	}

	return results
}
