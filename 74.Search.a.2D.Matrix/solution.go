package main

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix)
	for i < j {
		mid := i + (j-i)/2
		if matrix[mid][0] <= target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	row := i - 1
	if row < 0 {
		return false
	}
	i, j = 0, len(matrix[row])
	for i < j {
		mid := i + (j-i)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return false
}
