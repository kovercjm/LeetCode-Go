package main

func reverseWords(s string) string {
	runes := []rune(s)
	for i, next := 0, 1; i < len(runes); i, next = next+1, i+1 {
		for next < len(runes) && runes[next] != ' ' {
			next++
		}
		for j := next - 1; j > i; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}
	return string(runes)
}
