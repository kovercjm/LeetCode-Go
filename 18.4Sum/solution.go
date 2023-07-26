package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for k := 0; k < len(nums)-3; {
		for l := k + 1; l < len(nums)-2; {
			if nums[k]+nums[l] >= target && nums[l] > 0 {
				break
			}
			i, j := l+1, len(nums)-1
			for i < j {
				for i < j && nums[k]+nums[l]+nums[i]+nums[j] > target {
					j--
				}
				for i < j && nums[k]+nums[l]+nums[i]+nums[j] < target {
					i++
				}
				if i < j && nums[k]+nums[l]+nums[i]+nums[j] == target {
					result = append(result, []int{nums[k], nums[l], nums[i], nums[j]})
					for i++; i < j && nums[i] == nums[i-1]; i++ {
					}
					for j--; i < j && nums[j] == nums[j+1]; j-- {
					}
				}
			}
			for l++; l < len(nums)-2 && nums[l] == nums[l-1]; l++ {
			}
		}
		for k++; k < len(nums)-3 && nums[k] == nums[k-1]; k++ {
		}
	}
	return result
}
