package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			if _, ok := sum[num1+num2]; !ok {
				sum[num1+num2] = 1
			} else {
				sum[num1+num2]++
			}
		}
	}
	var result int
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			if count, ok := sum[-num3-num4]; ok {
				result += count
			}
		}
	}
	return result
}
