// https://leetcode.com/problems/jump-game/

package islandperimeter

func Input() [][]int {
	return [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
}

func Run(nums [][]int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if nums[i][j] == 0 {
				continue
			}
			ret += 4
			if j+1 < len(nums[i]) && nums[i][j+1] == 1 {
				ret--
			}
			if j > 0 && nums[i][j-1] == 1 {
				ret--
			}
			if i+1 < len(nums) && nums[i+1][j] == 1 {
				ret--
			}
			if i > 0 && nums[i-1][j] == 1 {
				ret--
			}
		}
	}
	return ret
}
