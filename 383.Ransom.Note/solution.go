package main

func canConstruct(ransomNote string, magazine string) bool {
	var count [26]int
	for _, char := range magazine {
		count[char-'a']++
	}
	for _, char := range ransomNote {
		count[char-'a']--
		if count[char-'a'] < 0 {
			return false
		}
	}
	return true
}
