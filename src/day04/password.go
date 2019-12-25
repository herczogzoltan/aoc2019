package day04

import (
	"strconv"
)

func FindPasswords(f int, t int) int {
	p := 0
	for i := f; i <= t; i++ {
		if isPassword(strconv.Itoa(i)) {
			p++
		}
	}

	return p
}

func isPassword(p string) bool {
	as := false
	for i := 0; i < len(p)-1; i++ {
		c, _ := strconv.Atoi(string(p[i]))
		n, _ := strconv.Atoi(string(p[i+1]))

		if c > n {
			return false
		}

		if !as {
			if c == n {
				as = true
			}
		}
	}

	return as
}
