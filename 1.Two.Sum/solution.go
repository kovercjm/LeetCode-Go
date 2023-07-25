package main

func twoSum(nums []int, target int) []int {
	positions := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := positions[target-nums[i]]; ok {
			return []int{i, j}
		}
		positions[nums[i]] = i
	}
	return nil
}
