package main

//func isAnagram(s string, t string) bool {
//	charMap := make(map[string]int)
//	for _, char := range s {
//		if _, ok := charMap[string(char)]; !ok {
//			charMap[string(char)] = 0
//		}
//		charMap[string(char)]++
//	}
//	for _, char := range t {
//		if _, ok := charMap[string(char)]; !ok {
//			return false
//		}
//		charMap[string(char)]--
//	}
//	for char := range charMap {
//		if charMap[char] != 0 {
//			return false
//		}
//	}
//	return true
//}
