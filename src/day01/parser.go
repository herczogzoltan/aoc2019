package day01

import (
	"bytes"
	"strconv"
)

func ParseInput(in []byte) []int64  {
	results := make([]int64, 0)
	lines := bytes.Split(in, []byte("\n"))
	for _, v := range lines {
		i, _ := strconv.Atoi(string(v))
		results = append(results, int64(i))
	}
	return results
}
