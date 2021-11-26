// https://leetcode.com/problems/search-insert-position/

package searchinsertposition

func Input() ([]int, int) {
	return []int{1, 3, 5, 6}, 7
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := (low + high) >> 1
		if target < nums[mid] {
			high = mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}

func Run(nums []int, target int) int {
	return searchInsert(nums, target)
}
