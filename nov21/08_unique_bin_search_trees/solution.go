// https://leetcode.com/problems/unique-binary-search-trees/

package uniquebinsearchtrees

func Input() int {
	return 3
}

func numTrees(n int) int {
	if n <= 1 {
		return n
	}

	C := make([]int, n+1)
	C[0], C[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i/2; j++ {
			C[i] += C[j] * C[i-j-1]
		}

		C[i] *= 2
		if i&1 == 1 {
			C[i] += C[i/2] * C[i/2]
		}
	}
	return C[n]
}

func Run(n int) int {
	return numTrees(n)
}
