package main

func reverseWords(s string) string {
	runes := []rune(s)
	if len(runes) <= 1 {
		return s
	}

	var fast, slow int
	for fast < len(runes) && runes[fast] == ' ' {
		fast++
	}
	for ; fast < len(runes); fast++ {
		if runes[fast] == ' ' && runes[fast-1] == ' ' {
			continue
		}
		runes[slow] = runes[fast]
		slow++
	}
	if runes[slow-1] == ' ' {
		slow--
	}

	reverse(runes, 0, slow-1)
	for i, j := 0, 0; i < slow; i, j = j+1, j+1 {
		for j < slow && runes[j] != ' ' {
			j++
		}
		reverse(runes, i, j-1)
	}
	return string(runes[:slow])
}

func reverse(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}
