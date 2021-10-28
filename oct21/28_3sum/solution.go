// https://leetcode.com/problems/3sum/

package threesum

import (
	"sort"
)

func Input() []int {
	return []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
}

func Run(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	N := len(nums)
	for i := 0; i < N-2; i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, N-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// skip all duplicates
				for l < r && nums[l+1] == nums[l] {
					l++
				}
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

// func Run(nums []int) [][]int {
//   sort.Ints(nums)
//   res := make([][]int, 0)
//   for i := 0; i < len(nums)-2; i++ {
//     if nums[i] > 0 {
//       return res
//     }
//     if i > 0 && nums[i] == nums[i-1] {
//       continue
//     }
//     f := false
//     for j := i + 1; j < len(nums)-1; j++ {
//       if f && nums[j] == nums[j-1] {
//         continue
//       }
//       if nums[i]+nums[j] > 0 {
//         break
//       }
//       for k := j + 1; k < len(nums); k++ {
//         if nums[i]+nums[j]+nums[k] == 0 {
//           res = append(res, []int{nums[i], nums[j], nums[k]})
//           break
//         }
//       }
//       f = true
//     }
//   }
//   return res
// }
