package day02

func IntCode(a []int) int {
	i := 0
	for i < len(a) {
		opCode := a[i]

		switch opCode {
		case 1:
			a[a[i+3]] = a[a[i+1]] + a[a[i+2]]
		case 2:
			a[a[i+3]] = a[a[i+1]] * a[a[i+2]]
		case 99:
			break
		}

		i += 4
	}

	return a[0]
}
