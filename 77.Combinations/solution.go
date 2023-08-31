package main

func combine(n int, k int) [][]int {
	return subCombine([]int{}, 1, n, k)
}

func subCombine(prefix []int, start, n, k int) [][]int {
	var result [][]int
	for i := start; i <= n-(k-len(prefix)-1); i++ {
		newPrefix := append(prefix, i)
		if len(prefix) == k-1 {
			var res []int
			res = append(res, prefix...)
			result = append(result, append(res, i))
			continue
		}
		result = append(result, subCombine(newPrefix, i+1, n, k)...)
	}
	return result
}
