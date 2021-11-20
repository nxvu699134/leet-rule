// https://leetcode.com/problems/single-element-in-a-sorted-array/

package singleelinsortedarray

import "fmt"

func Input() []int {
	return []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	// return []int{3, 3, 7, 7, 10, 11, 11}
}

// func singleNonDuplicate(nums []int) int {
//   if len(nums) == 1 {
//     return nums[0]
//   }
//
//   if nums[0] != nums[1] {
//     return nums[0]
//   }
//
//   if nums[len(nums)-1] != nums[len(nums)-2] {
//     return nums[len(nums)-1]
//   }
//
//   low, high := 0, len(nums)-1
//   for low < high {
//     mid := (low + high) >> 1
//     if nums[mid] == nums[mid-1] {
//       if mid&1 == 1 {
//         low = mid - 1
//       } else {
//         high = mid
//       }
//     } else if nums[mid] == nums[mid+1] {
//       if mid&1 == 1 {
//         high = mid + 1
//       } else {
//         low = mid
//       }
//     } else {
//       return nums[mid]
//     }
//   }
//   return 0
// }

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) >> 1
		fmt.Println(low, mid, high)
		var consider int
		if mid&1 == 1 {
			consider = mid - 1
		} else {
			consider = mid + 1
		}
		if nums[mid] == nums[consider] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
func Run(nums []int) int {
	return singleNonDuplicate(nums)
}
