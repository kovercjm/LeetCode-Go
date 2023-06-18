package main

func commonChars(words []string) []string {
	var (
		alphabets [26]int
		result    []string
	)

	for i, num := range count(words[0]) {
		alphabets[i] += num
	}

	for i := 1; i < len(words); i++ {
		for i, num := range count(words[i]) {
			if num < alphabets[i] {
				alphabets[i] = num
			}
		}
	}

	for alphabet, c := range alphabets {
		for i := 1; i <= c; i++ {
			result = append(result, string(rune('a'+alphabet)))
		}
	}

	return result
}

func count(word string) (result [26]int) {
	for _, char := range word {
		result[char-'a']++
	}
	return
}
