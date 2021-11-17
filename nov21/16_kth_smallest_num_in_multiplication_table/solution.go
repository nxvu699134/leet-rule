// https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/

package kthsmallestnuminmultiplicationtable

import (
	"fmt"
)

func Input() (int, int, int) {
	return 2, 3, 6
}

//Naive solution
// func findKthNumber(m int, n int, k int) int {
//   arr := make([]int, m*n)
//   for i := 1; i <= m; i++ {
//     for j := 1; j <= n; j++ {
//       arr[(i-1)*n+j-1] = i * j
//     }
//   }
//   sort.Ints(arr)
//   return arr[k-1]
// }

func printTable(m, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%2d ", i*j)
		}
		fmt.Println()
	}
}

func fromMatrix(m, n int) func(int) int {
	return func(val int) int {
		ret := 0
		for i := 1; i <= m; i++ {
			if numBefore := val / i; numBefore > n {
				ret += n
			} else {
				ret += numBefore
			}
		}
		return ret
	}
}

func findKthNumber(m int, n int, k int) int {
	if m > n {
		m, n = n, m
	}
	low, high := 1, m*n
	ret := 0
	checkNumBefore := fromMatrix(m, n)
	for low <= high {
		mid := (low + high) / 2
		if c := checkNumBefore(mid); c < k {
			low = mid + 1
		} else {
			high = mid - 1
			ret = mid
		}

	}
	return ret
}

func Run(m, n, k int) int {
	return findKthNumber(m, n, k)
}
