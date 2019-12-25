package day02

import (
	"bytes"
	"strconv"
)

func Parser(input []uint8) []int {
	bytesArr := bytes.Split(input, []byte(","))
	numbers := make([]int, len(bytesArr))

	for i := 0; i < len(bytesArr); i++ {
		n, _ := strconv.Atoi(string(bytesArr[i]))
		numbers[i] = n
	}

	return numbers
}
