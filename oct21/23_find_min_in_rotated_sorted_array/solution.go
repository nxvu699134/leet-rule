// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

package findmininrotatedsortedarray

func Input() []int {
	return []int{4, 5, 6, 7, 0, 1, 4}
}

// This solution better than binary search aspect of space and time complexity
func Run(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			return nums[i]
		}
	}
	return nums[0]
}

// func Run(nums []int) int {
//   l, r := 0, len(nums)-1
//   for l < r {
//     m := (l + r) / 2
//     if nums[m] > nums[r] {
//       l = m + 1
//     } else if nums[m] < nums[r] {
//       r = m
//     } else {
//       r--
//     }
//   }
//   return nums[r]
// }
