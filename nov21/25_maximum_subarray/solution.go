// https://leetcode.com/problems/maximum-subarray/

package maximumsubarray

func Input() []int {
	return []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
}

// Naive solution
// O(n^2)
// func maxSubArray(nums []int) int {
// const MIN_INT = -1 << 31
// max := MIN_INT
// for i := 0; i < len(nums); i++ {
//   sum := 0
//   for j := i; j < len(nums); j++ {
//     sum += nums[j]
//     if sum > max {
//       max = sum
//     }
//   }
// }
// return max
// }

const MIN_INT = -1 << 31

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// Recursion solution
// O(n^2)
// func _maxSubArray(nums []int, i int, isInclude bool) int {
//   if i == len(nums) {
//     if isInclude { // have include here thats mean the prev call is the last element => so just stop
//       return 0
//     }
//     return MIN_INT
//   }
//   if isInclude {
//     // 0 because in callstack at i + 1 will return 0 thats mean the return become nums[i] + 0 => so we stop here dont add the next part
//     return max(0, nums[i]+_maxSubArray(nums, i+1, true))
//   }
//   return max(_maxSubArray(nums, i+1, false), nums[i]+_maxSubArray(nums, i+1, true))
// }
//
// func maxSubArray(nums []int) int {
//   return _maxSubArray(nums, 0, false)
// }

// Do some optimizes for the above solution with dy-pro
// func _maxSubArray(nums []int, i int, isInclude bool, dp map[bool][]int) int {
//   if i == len(nums) {
//     if isInclude { // have include here thats mean the prev call is the last element => so just stop
//       return 0
//     }
//     return MIN_INT
//   }
//   if dp[isInclude][i] != MIN_INT {
//     return dp[isInclude][i]
//   }
//
//   if isInclude {
//     // 0 because in callstack at i + 1 will return 0 thats mean the return become nums[i] + 0 => so we stop here dont add the next part
//     dp[true][i] = max(0, nums[i]+_maxSubArray(nums, i+1, true, dp))
//     return dp[true][i]
//   }
//   dp[false][i] = max(_maxSubArray(nums, i+1, false, dp), nums[i]+_maxSubArray(nums, i+1, true, dp))
//   return dp[false][i]
// }
//
// func maxSubArray(nums []int) int {
//   dp := make(map[bool][]int)
//   dp[false] = make([]int, len(nums))
//   dp[true] = make([]int, len(nums))
//   for i := 0; i < len(nums); i++ {
//     dp[false][i] = MIN_INT
//     dp[true][i] = MIN_INT
//   }
//   return _maxSubArray(nums, 0, false, dp)
// }

// func maxSubArray(nums []int) int {
//   dp := make(map[bool][]int)
//   dp[false] = make([]int, len(nums))
//   dp[true] = make([]int, len(nums))
//   for i := 0; i < len(nums); i++ {
//     dp[false][i] = MIN_INT
//     dp[true][i] = MIN_INT
//   }
//   //We can re write recursion with a simple for
//   // Notice the next result is calculated from prev
//   // Notice What is return if i = 0 in recursion so we can get the init value
//   dp[false][0], dp[true][0] = nums[0], nums[0]
//   for i := 1; i < len(nums); i++ {
//     dp[true][i] = max(nums[i], dp[true][i-1]+nums[i])
//     dp[false][i] = max(dp[true][i], dp[false][i-1])
//   }
//   return dp[false][len(nums)-1]
// }

func maxSubArray(nums []int) int {
	maxInclueOrExclude, maxInclude := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxInclude = max(nums[i], maxInclude+nums[i])
		maxInclueOrExclude = max(maxInclude, maxInclueOrExclude)
	}
	return maxInclueOrExclude
}

func Run(nums []int) int {
	return maxSubArray(nums)
}
