// https://leetcode.com/problems/dungeon-game/
package dungeongame

func Input() [][]int {
	// return [][]int{{-200}}
	return [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
}

func abs(n int) int {
	mask := n >> 31
	return (n ^ mask) - mask
}

//
// func travel(grid [][]int, i, j, cur_point, start_point int, min *int) {
//   p := grid[i][j]
//   t := cur_point + p
//   if t <= 0 {
//     cur_point = 1
//     start_point = start_point + abs(t) + 1
//   } else {
//     cur_point = t
//   }
//
//   // fmt.Printf("%d %d, %d %d\n", i, j, cur_point, start_point)
//   if i == len(grid)-1 && j == len(grid[0])-1 {
//     if start_point < *min {
//       *min = start_point
//     }
//     return
//   }
//   if i+1 < len(grid) {
//     travel(grid, i+1, j, cur_point, start_point, min)
//   }
//
//   if j+1 < len(grid[0]) {
//     travel(grid, i, j+1, cur_point, start_point, min)
//   }
// }
//
// func Run(grid [][]int) int {
//   min := 1 << 32
//   travel(grid, 0, 0, 1, 1, &min)
//   return min
// }

func Run(grid [][]int) int {

	Abs := func(n int) int {
		mask := n >> 31
		return (n ^ mask) - mask
	}

	GetValue := func(i, j int) int {
		if i == len(grid) || j == len(grid[0]) {
			return (1 << 31) * -1
		}
		return grid[i][j]
	}

	Max := func(n1, n2 int) int {
		if n1 < n2 {
			return n2
		}
		return n1
	}

	nRow, nCol := len(grid), len(grid[0])

	if grid[nRow-1][nCol-1] > 0 {
		grid[nRow-1][nCol-1] = 0
	}

	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if i == len(grid)-1 && j == len(grid[0])-1 {
				continue
			}
			health_to_lose := Max(GetValue(i+1, j), GetValue(i, j+1)) + grid[i][j]
			// fmt.Println(health_to_lose)
			if health_to_lose > 0 {
				grid[i][j] = 0
			} else {
				grid[i][j] = health_to_lose
			}
		}
	}
	return Abs(grid[0][0]) + 1
}
