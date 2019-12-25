package day04

import (
	"bytes"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func ParseInput(in []byte) (int, int) {
	numbers := bytes.Split(in, []byte("-"))

	if len(numbers) > 2 {
		log.Error("There are more than 2 numbers as interval")
	}

	f, err := strconv.Atoi(string(numbers[0]))

	if err != nil {
		log.Error(err)
	}

	t, err := strconv.Atoi(string(numbers[1]))

	if err != nil {
		log.Error(err)
	}

	return f, t
}
