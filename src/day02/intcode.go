package day02

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	ADDITION = iota + 1
	MULTIPLY
	STORE
	OUTPUT
	QUIT = 99
)

func IntCode(a map[int]int) int {
	i := 0
	params := 4
	for i < len(a) {
		opCodeWithModes := a[i]
		opCode, err := getOpCode(opCodeWithModes)
		if err != nil {
			log.Error(err)
		}
		modes, err := getModes(opCodeWithModes)
		if err != nil {
			log.Error(err)
		}

		switch opCode {
		case ADDITION:
			args := getArguments(3, modes, i, a)
			a[args[2]] = args[0] + args[1]
			params = 4
		case MULTIPLY:
			args := getArguments(3, modes, i, a)
			a[args[2]] = args[0] * args[1]
			params = 4
		case STORE:
			a[a[i+1]] = 1
			params = 2
		case OUTPUT:
			fmt.Printf("OpCode 4: %v\n", a[a[i+1]])
			params = 2
		case QUIT:
			break
		}

		i += params
	}

	return a[0]
}

func getArguments(argCount int, modes []int, pos int, addresses map[int]int) (arguments []int) {
	for i := 0; i < argCount; i++ {
		var mode int
		if i >= len(modes) {
			mode = 0
		} else {
			mode = modes[i]
		}

		switch mode {
		case 0:
			if i+1 == argCount {
				arguments = append(arguments, addresses[pos+i+1])
			} else {
				arguments = append(arguments, addresses[addresses[pos+i+1]])
			}
		case 1:
			arguments = append(arguments, addresses[pos+i+1])
		}

	}

	return
}

func getOpCode(opCodeWithModes int) (opCode int, err error) {
	sopCodeWithModes := strconv.Itoa(opCodeWithModes)
	l := len(sopCodeWithModes)
	if l == 1 {
		return opCodeWithModes, nil
	}

	opCode, err = strconv.Atoi(sopCodeWithModes[l-2:])

	return
}

func getModes(opCodeWithModes int) (modes []int, err error) {
	opCodeWithModesStr := strconv.Itoa(opCodeWithModes)
	l := len(opCodeWithModesStr)
	if l == 1 {
		return
	}

	modesStr := opCodeWithModesStr[:l-2]

	for i := len(modesStr) - 1; i >= 0; i-- {
		m, _ := strconv.Atoi(string(modesStr[i]))
		modes = append(modes, m)
	}

	return
}
