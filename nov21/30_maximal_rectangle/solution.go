// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/

package minvaluetogetpositive

func Input() [][]byte {
	return [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	// return [][]byte{
	//   {'0', '1'},
	//   {'1', '0'},
	// }
}

// Naive solution
// func maxAreaFrom(matrix [][]byte, row, col int) int {
//   minConsecutive1s := len(matrix[0])
//   maxArea := 0
//   for i := row; i < len(matrix); i++ {
//     consecutive1sInRow := 0
//     for j := col; j < len(matrix[i]); j++ {
//       if matrix[i][j] == '1' {
//         consecutive1sInRow++
//       } else {
//         break
//       }
//     }
//
//     if consecutive1sInRow == 0 {
//       return maxArea
//     }
//
//     if consecutive1sInRow < minConsecutive1s {
//       minConsecutive1s = consecutive1sInRow
//     }
//
//     if area := minConsecutive1s * (i - row + 1); area > maxArea {
//       maxArea = area
//     }
//   }
//   return maxArea
// }
//
// func maximalRectangle(matrix [][]byte) int {
//   maxArea := 0
//   for i := 0; i < len(matrix); i++ {
//     for j := 0; j < len(matrix[i]); j++ {
//       if area := maxAreaFrom(matrix, i, j); area > maxArea {
//         maxArea = area
//       }
//     }
//   }
//   return maxArea
// }

func maxAreaFrom(row, col int, hist [][]int) int {
	minConsecutive1s := len(hist[0])
	maxArea := 0
	for i := row; i < len(hist); i++ {
		if hist[i][col] == 0 {
			return maxArea
		}

		if hist[i][col] < minConsecutive1s {
			minConsecutive1s = hist[i][col]
		}

		if area := minConsecutive1s * (i - row + 1); area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func maximalRectangle(matrix [][]byte) int {
	// use Histogram of 1s in row instead recalculate from each i, j
	hist := make([][]int, len(matrix))
	for i := 0; i < len(hist); i++ {
		hist[i] = make([]int, len(matrix[i]))
		hist[i][len(matrix[i])-1] = int(matrix[i][len(matrix[i])-1] - '0')
		for j := len(matrix[i]) - 2; j >= 0; j-- {
			if matrix[i][j] == '0' {
				continue
			}
			hist[i][j] = hist[i][j+1] + 1
		}
	}
	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if area := maxAreaFrom(i, j, hist); area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

//Another solution is use monotonic stack
// Implement latter

func Run(matrix [][]byte) int {
	return maximalRectangle(matrix)
}
