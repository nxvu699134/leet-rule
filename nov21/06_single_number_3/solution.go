// https://leetcode.com/problems/single-number-iii/

package singlenumber3

import (
	"sort"
)

func Input() []int {
	return []int{1, 2, 1, 3, 2, 5}
}

func singleNumber(nums []int) []int {
	sort.Ints(nums)
	res := make([]int, 0, 2)
	i := 0
	for i < len(nums)-1 {
		if nums[i+1] == nums[i] {
			i += 2
		} else {
			res = append(res, nums[i])
			i++
		}
	}
	if i == len(nums)-1 {
		res = append(res, nums[len(nums)-1])
	}
	return res
}

func Run(nums []int) []int {
	return singleNumber(nums)
}
