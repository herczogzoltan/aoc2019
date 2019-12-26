package day04

import (
	"strconv"
)

type passwordChecker func(string) bool

func FindPasswords(f int, t int, c passwordChecker) int {
	p := 0
	for i := f; i <= t; i++ {
		if c(strconv.Itoa(i)) {
			p++
		}
	}

	return p
}

func IsPassword(p string) bool {
	as := false
	for i := 0; i < len(p)-1; i++ {
		c, _ := strconv.Atoi(string(p[i]))
		n, _ := strconv.Atoi(string(p[i+1]))

		if c > n {
			return false
		}

		if c == n {
			as = true
		}
	}

	return as
}

func IsStrongerPassword(p string) bool {
	as := false
	for i := 0; i < len(p)-1; i++ {
		c, _ := strconv.Atoi(string(p[i]))
		n, _ := strconv.Atoi(string(p[i+1]))

		if c > n {
			return false
		}

		if c == n {
			if (i+2 < len(p) && p[i+2] != p[i+1]) || (i+2 >= len(p)) {
				as = true
			}

			for j := i + 2; j < len(p); j++ {
				if p[j] == p[i] {
					i++
					continue
				}
				break
			}
		}
	}

	return as
}
