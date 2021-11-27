// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/

package minvaluetogetpositive

func Input() []int {
	return []int{1, 2, 3, 4}
}

func productExceptSelf(nums []int) []int {
	pre, suf := 1, 1
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = pre
		pre *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= suf
		suf *= nums[i]
	}
	return ret
}

func Run(nums []int) []int {
	return productExceptSelf(nums)
}
