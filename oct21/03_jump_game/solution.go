// https://leetcode.com/problems/jump-game/

package jumpgame

func Input() []int {
	// return []int{2,3,1,1,4}
	return []int{3, 2, 1, 0, 4}
}

func Run(nums []int) bool {
	n := len(nums)
	if 1 == n {
		return true
	}

	by_pass_zero := 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < by_pass_zero {
			by_pass_zero++
		} else {
			by_pass_zero = 1
		}
	}
	return by_pass_zero == 1
}
