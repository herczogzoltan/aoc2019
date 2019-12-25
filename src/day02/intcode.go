package day02

func IntCode(a []int, noun int, verb int) int {
	i := 0
	// because 1202 program alarm: https://www.hq.nasa.gov/alsj/a11/a11.landing.html#1023832
	a[1] = noun
	a[2] = verb

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
