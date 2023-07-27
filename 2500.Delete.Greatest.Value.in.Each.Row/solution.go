package main

import "sort"

func deleteGreatestValue(grid [][]int) int {
	for _, row := range grid {
		sort.Ints(row)
	}
	var result int
	for i := 0; i < len(grid[0]); i++ {
		var max int
		for j := 0; j < len(grid); j++ {
			if max < grid[j][i] {
				max = grid[j][i]
			}
		}
		result += max
	}
	return result
}
