package main

func sortedSquares(nums []int) []int {
	result, k := make([]int, len(nums)), len(nums)-1
	for i, j := 0, len(nums)-1; i <= j && k >= 0; k-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[k] = nums[i] * nums[i]
			i++
		} else {
			result[k] = nums[j] * nums[j]
			j--
		}
	}
	return result
}
