// https://leetcode.com/problems/unique-paths-iii/

package uniquepath3

import "fmt"

func Input() [][]int {
	return [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}
}

func printGrid(grid [][]int) {
	fmt.Println()
	for i := range grid {
		for j := range grid[0] {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

}

const (
	START = 1
	END   = 2
	WALL  = -1
	EMPTY = 0
)

func Dfs(grid [][]int, row int, col int, cur_step int, max_step int) int {
	if grid[row][col] == END {
		if cur_step == max_step {
			return 1
		} else {
			return 0
		}
	}
	res := 0
	tmp := grid[row][col]
	grid[row][col] = WALL // treat visited cell as wall

	if left_col := col - 1; left_col >= 0 && grid[row][left_col] != WALL {
		res += Dfs(grid, row, left_col, cur_step+1, max_step)
	}
	if right_col := col + 1; right_col < len(grid[0]) && grid[row][right_col] != WALL {
		res += Dfs(grid, row, right_col, cur_step+1, max_step)
	}
	if under_row := row + 1; under_row < len(grid) && grid[under_row][col] != WALL {
		res += Dfs(grid, under_row, col, cur_step+1, max_step)
	}
	if upper_row := row - 1; upper_row >= 0 && grid[upper_row][col] != WALL {
		res += Dfs(grid, upper_row, col, cur_step+1, max_step)
	}

	grid[row][col] = tmp //break visited wall
	return res
}

func uniquePathsIII(grid [][]int) int {
	s_row, s_col := - -1, -1
	max_step := 2 // add start and end point
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == START {
				s_row, s_col = i, j
			}
			if grid[i][j] == EMPTY {
				max_step++
			}
		}
	}
	return Dfs(grid, s_row, s_col, 1, max_step)
}

func Run(grid [][]int) int {
	return uniquePathsIII(grid)
}
