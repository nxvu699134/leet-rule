// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/

package minvaluetogetpositive

func Input() []int {
	return []int{-3, 2, -3, 4, 2}
}

func minStartValue(nums []int) int {
	ret, cur := 1, 1
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		if cur < 1 {
			ret += -cur + 1
			cur = 1
		}
	}
	return ret
}

func Run(nums []int) int {
	return minStartValue(nums)
}
