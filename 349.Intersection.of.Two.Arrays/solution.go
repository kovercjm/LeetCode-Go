package main

func intersection(nums1 []int, nums2 []int) []int {
	exist := make(map[int]struct{})
	for _, num := range nums1 {
		if _, ok := exist[num]; !ok {
			exist[num] = struct{}{}
		}
	}
	var res []int
	for _, num := range nums2 {
		if _, ok := exist[num]; ok {
			res = append(res, num)
			delete(exist, num)
		}
	}
	return res
}
