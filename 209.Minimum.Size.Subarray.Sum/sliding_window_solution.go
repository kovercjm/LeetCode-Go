package main

//func minSubArrayLen(target int, nums []int) int {
//	i, j, sum, length := 0, 0, 0, len(nums)+1
//	for {
//		for ; sum < target && j < len(nums); j++ {
//			sum += nums[j]
//		}
//		if j == len(nums) && sum < target {
//			break
//		}
//
//		for ; sum >= target && i < j; i++ {
//			sum -= nums[i]
//		}
//		if j-i+1 < length {
//			length = j - i + 1
//		}
//	}
//	if length == len(nums)+1 {
//		return 0
//	}
//	return length
//}
