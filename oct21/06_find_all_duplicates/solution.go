// https://leetcode.com/problems/find-all-duplicates-in-an-array/

package findallduplicates

// import "sort"

func Input() []int {
	// return []int{4, 3, 2, 7, 8, 2, 3, 1}
	return []int{2, 2}
}

func Run(nums []int) []int {
	ret := make([]int, 0)
	l := len(nums)
	for _, v := range nums {
		point_to := v%(l+1) - 1
		point_value := nums[point_to]
		if point_value > l {
			ret = append(ret, point_to+1)
		} else {
			nums[point_to] = point_value + l + 1
		}
	}
	return ret
}

// O(nlogn)
// func Run(nums []int) []int {
//   sort.Ints(nums)
//   ret := make([]int, 0)
//   f := false
//   for i := 1; i < len(nums); i++ {
//     if nums[i] == nums[i-1] {
//       if f == true {
//         continue
//       } else {
//         ret = append(ret, nums[i-1])
//       }
//     } else {
//       f = false
//     }
//   }
//   _ = f
//
//   return ret
// }
