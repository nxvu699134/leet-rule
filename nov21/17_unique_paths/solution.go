// https://leetcode.com/problems/unique-paths/

package uniquepaths

func Input() (int, int) {
	return 3, 7
}

// func fromMatrix(m, n int) func(int, int) int {
//   var Dfs func(int, int) int
//   Dfs = func(row, col int) int {
//     if row == m-1 && col == n-1 {
//       return 1
//     }
//     ret := 0
//     if row+1 < m {
//       ret += Dfs(row+1, col)
//     }
//     if col+1 < n {
//       ret += Dfs(row, col+1)
//     }
//     return ret
//   }
//   return Dfs
// }
//
// func uniquePaths(m int, n int) int {
//   Dfs := fromMatrix(m, n)
//   return Dfs(0, 0)
// }

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func Run(m, n int) int {
	return uniquePaths(m, n)
}
