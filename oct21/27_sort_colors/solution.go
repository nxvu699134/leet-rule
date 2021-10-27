// https://leetcode.com/problems/sort-colors/

package sortcolors

func Input() []int {
	return []int{2, 0, 2, 1, 1, 0}
}

// This solution better than binary search aspect of space and time complexity
// func Run(nums []int) []int {
//   for i := 0; i < len(nums)-1; i++ {
//     for j := i + 1; j < len(nums); j++ {
//       if nums[j] < nums[i] {
//         nums[i], nums[j] = nums[j], nums[i]
//       }
//     }
//   }
//   return nums
// }

func Run(nums []int) []int {
	// Dutch national flag problem
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[mid], nums[low] = nums[low], nums[mid]
			mid++
			low++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
	return nums
}
