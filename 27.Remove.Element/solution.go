package main

func removeElement(nums []int, val int) int {
	end := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			end++
			nums[end] = nums[i]
		}
	}
	return end + 1
}
