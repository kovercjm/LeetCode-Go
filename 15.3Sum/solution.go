package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for k := 0; k < len(nums)-2; {
		if nums[k] > 0 {
			break
		}
		i, j := k+1, len(nums)-1
		for i < j {
			for nums[i]+nums[j]+nums[k] > 0 && i < j {
				j--
			}
			for nums[i]+nums[j]+nums[k] < 0 && i < j {
				i++
			}
			if i < j && nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for i++; i < j && nums[i] == nums[i-1]; i++ {
				}
				for j--; i < j && nums[j] == nums[j+1]; j-- {
				}
			}
		}
		for k++; k < len(nums)-2 && nums[k] == nums[k-1]; k++ {
		}
	}

	return result
}
