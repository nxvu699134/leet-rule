// https://leetcode.com/problems/house-robber/

package houserobber

func Input() []int {
	return []int{1, 2, 3, 1}
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func rob(nums []int) int {
	rob, wait := nums[0], 0
	for i := 1; i < len(nums); i++ {
		last_rob := rob
		rob = max(rob, wait+nums[i])
		wait = max(wait, last_rob)
	}
	if rob > wait {
		return rob
	}
	return wait
}

func Run(nums []int) int {
	return rob(nums)
}
