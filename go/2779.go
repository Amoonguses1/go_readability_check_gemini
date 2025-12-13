package main

import "sort"

// Returns the maximum length of a subsequence
// where all elements can be made equal by modifying each element
// at most once within the range [nums[i] - k, nums[i] + k].
func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)

	start := 0
	for _, num := range nums {
		if num > nums[start]+2*k {
			start++
		}
	}
	return len(nums) - start
}
