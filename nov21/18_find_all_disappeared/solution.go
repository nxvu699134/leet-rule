// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

package findalldisappeared

func Input() []int {
	return []int{4, 3, 2, 7, 8, 2, 3, 1}
}

// func findDisappearedNumbers(nums []int) []int {
//   ret := make([]int, 0)
//   for i := 0; i < len(nums); i++ {
//     point_to := nums[i]
//     if point_to < 0 {
//       point_to *= -1
//     }
//     point_to--
//     if nums[point_to] > 0 {
//       nums[point_to] *= -1
//     }
//   }
//   for i := 0; i < len(nums); i++ {
//     if nums[i] > 0 {
//       ret = append(ret, i+1)
//     }
//   }
//   return ret
// }

func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, 0)
	const mask = 1 << 16
	for i := 0; i < len(nums); i++ {
		nums[nums[i] & ^mask - 1] |= mask
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]&mask == 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func Run(nums []int) []int {
	return findDisappearedNumbers(nums)
}
