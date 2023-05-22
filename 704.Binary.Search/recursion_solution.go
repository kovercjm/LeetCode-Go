package main

//func search(nums []int, target int) int {
//	if len(nums) == 0 {
//		return -1
//	}
//	mid := len(nums) / 2
//	if nums[mid] == target {
//		return mid
//	} else if nums[mid] < target {
//		result := search(nums[mid+1:], target)
//		if result == -1 {
//			return -1
//		}
//		return mid + 1 + result
//	} else {
//		return search(nums[:mid], target)
//	}
//}
