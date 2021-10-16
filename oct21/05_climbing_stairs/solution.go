// https://leetcode.com/problems/climbing-stairs/

package climbingstairs

// import "fmt"

func Input() int {
	return 45
}

// func combinations(n int, k int) int {
//   if n == k {
//     return 1
//   }
//   if k == 1 {
//     return n
//   }
//
//   var cache_min, cache_max int
//   if k < n-k {
//     cache_min, cache_max = k, n-k
//   } else {
//     cache_min, cache_max = n-k, k
//   }
//
//   f_min := 1
//   for i := 1; i <= cache_min; i++ {
//     f_min *= i
//   }
//
//   f := uint64(n)
//   for i := n - 1; i > cache_max; i-- {
//     f *= uint64(i)
//   }
//
//   // fmt.Printf("%d, %d\n", cache_max, cache_min)
//   // fmt.Printf("%d, %d\n", f_min, f)
//   return int(f / uint64(f_min))
// }
//
// func Run(n int) int {
//   ret := 1
//   for i := 1; i <= n/2; i++ {
//     fmt.Printf("%d C %d = %d\n", n-i, i, combinations(n-i, i))
//     ret += int(combinations(n-i, i))
//   }
//   return ret
// }

func Run(n int) int {
	r, c := 1, 1
	for i := 2; i <= n; i++ {
		r, c = r+c, r
	}
	return r
}
