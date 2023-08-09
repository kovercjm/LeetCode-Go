package main

func subarraysDivByK(nums []int, k int) int {
	var (
		prefixSum []int
		count     = map[int]int{0: 1}
		s         int
	)
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		prefixSum = append(prefixSum, s)
	}

	for _, num := range prefixSum {
		if _, ok := count[(num%k+k)%k]; !ok {
			count[(num%k+k)%k] = 1
		} else {
			count[(num%k+k)%k]++
		}
	}
	s = 0
	for _, c := range count {
		s += c * (c - 1) / 2
	}
	return s
}
