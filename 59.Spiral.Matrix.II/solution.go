package main

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	direction, x, y := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		result[y][x] = i
		switch direction {
		case 0:
			if x+1 < n && result[y][x+1] == 0 {
				x++
				continue
			}
			direction = (direction + 1) % 4
			fallthrough
		case 1:
			if y+1 < n && result[y+1][x] == 0 {
				y++
				continue
			}
			direction = (direction + 1) % 4
			fallthrough
		case 2:
			if x-1 >= 0 && result[y][x-1] == 0 {
				x--
				continue
			}
			direction = (direction + 1) % 4
			fallthrough
		case 3:
			if y-1 >= 0 && result[y-1][x] == 0 {
				y--
				continue
			}
			direction = (direction + 1) % 4
			x++
		}
	}
	return result
}
