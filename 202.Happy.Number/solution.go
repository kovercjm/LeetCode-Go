package main

import (
	"math"
)

func isHappy(n int) bool {
	var (
		squares [10]int
		newNum  int
	)
	for i := 0; i <= 9; i++ {
		squares[i] = i * i
	}

	occured := make(map[int]struct{})
	occured[n] = struct{}{}
	for {
		for newNum = 0; n > 0; n /= 10 {
			newNum += squares[n%10]
		}
		if newNum == 1 || math.Log10(float64(newNum)) == 0 {
			return true
		}
		if _, ok := occured[newNum]; ok {
			return false
		}
		occured[newNum] = struct{}{}
		n = newNum
	}
}
