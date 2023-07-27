package main

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes); i += 2 * k {
		l, r := i, i+k-1
		if r >= len(bytes) {
			r = len(bytes) - 1
		}
		for ; l < r; l, r = l+1, r-1 {
			bytes[l], bytes[r] = bytes[r], bytes[l]
		}
	}
	return string(bytes)
}
