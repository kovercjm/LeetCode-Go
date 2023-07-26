package main

import "fmt"

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes) && (i+k-1 < len(bytes) || k > len(bytes)); i += 2 * k {
		l, r := i, i+k-1
		if r >= len(bytes) {
			r = len(bytes) - 1
		}
		if string(bytes[l]) == "l" {
			fmt.Println(l, r)
		}

		for ; l < r; l, r = l+1, r-1 {
			bytes[l], bytes[r] = bytes[r], bytes[l]
		}
	}
	return string(bytes)
}

func main() {
	reverseStr("limjkfnqcqnajmebeddqsgl", 39)
}
