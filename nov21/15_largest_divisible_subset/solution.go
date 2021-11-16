// https://leetcode.com/problems/largest-divisible-subset/

package largestdivisiblesubset

import (
	"fmt"
	"sort"
)

func Input() []int {
	return []int{1, 2, 4, 8}
}

// https://cp-algorithms.com/sequences/longest_increasing_subsequence.html
// based on this problem but modified condition
// [til] the ways to solve are so cool
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	prev := make([]int, len(nums)) // this array save the highest divisor of nums ith
	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[j] >= dp[i] {
				dp[i] = 1 + dp[j]
				prev[i] = j
			}
		}
	}

	max_idx := 0
	for i := 1; i < len(dp); i++ {
		if dp[i] > dp[max_idx] {
			max_idx = i
		}
	}
	ret := make([]int, dp[max_idx])
	j := dp[max_idx] - 1
	for max_idx != -1 {
		ret[j] = nums[max_idx]
		max_idx = prev[max_idx]
		j--
	}
	return ret
}

func Run(nums []int) []int {
	return largestDivisibleSubset(nums)
}
