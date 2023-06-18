package main

import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arrayS, arrayT := []byte(s), []byte(t)
	sort.Slice(arrayS, func(i, j int) bool { return arrayS[i] < arrayS[j] })
	sort.Slice(arrayT, func(i, j int) bool { return arrayT[i] < arrayT[j] })
	return string(arrayS) == string(arrayT)
}
