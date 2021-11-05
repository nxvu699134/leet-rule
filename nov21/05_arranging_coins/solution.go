// https://leetcode.com/problems/arranging-coins/

package arrangingcoins

import (
	"math"
)

func Input() int {
	return 5
}

// func arrangeCoins(n int) int {
//   l, r := 1, n
//   for l <= r {
//     m := (l + r) / 2
//     k := (m + 1) * m / 2
//     if k == n {
//       return m
//     } else if k < n {
//       l = m + 1
//     } else {
//       r = m - 1
//     }
//   }
//   return r
// }

func arrangeCoins(n int) int {
	return int(math.Sqrt(float64(2*n)+0.25) - 0.5)
}

func Run(n int) int {
	return arrangeCoins(n)
}
