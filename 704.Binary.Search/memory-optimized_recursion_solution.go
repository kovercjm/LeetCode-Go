package main

//func search(nums []int, target int) int {
//	array = nums
//	targetNum = target
//	return binarySearch(0, len(nums)-1)
//}

var (
	array     []int
	targetNum int
)

// Recursion binary search solution with memory-optimized
func binarySearch(left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2

	if array[mid] == targetNum {
		return mid
	} else if array[mid] < targetNum {
		return binarySearch(mid+1, right)
	} else {
		return binarySearch(left, mid-1)
	}
}
