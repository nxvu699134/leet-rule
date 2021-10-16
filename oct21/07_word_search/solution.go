// https://leetcode.com/problems/word-search/

package wordsearch

// import "fmt"

func Input() ([][]byte, string) {
	// return [][]byte{
	//   {'A', 'B', 'C', 'E'},
	//   {'S', 'F', 'C', 'S'},
	//   {'A', 'D', 'E', 'E'},
	// }, "SEE"
	return [][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
	}, "aaaaaaaaaaaaa"
	// return [][]byte{
	//     {'a', 'a', 'b', 'a', 'a', 'b'},
	//     {'b', 'a', 'b', 'a', 'b', 'b'},
	//     {'b', 'a', 'b', 'b', 'b', 'b'},
	//     {'a', 'a', 'b', 'a', 'b', 'a'},
	//     {'b', 'b', 'a', 'a', 'a', 'b'},
	//     {'b', 'b', 'b', 'a', 'b', 'a'},
	//   },
	//   "aaaababab"
}

//TODO: Do in someway thats not use recursion
func check(board [][]byte, word string, row, col int, depth int) bool {
	tmp := board[row][col]
	if board[row][col] == word[depth] {
		if depth == len(word)-1 {
			return true
		}
		board[row][col] = 0
		if row > 0 && board[row-1][col] != 0 && check(board, word, row-1, col, depth+1) {
			return true
		}
		if col > 0 && board[row][col-1] != 0 && check(board, word, row, col-1, depth+1) {
			return true
		}
		if row+1 < len(board) && board[row+1][col] != 0 && check(board, word, row+1, col, depth+1) {
			return true
		}
		if col+1 < len(board[0]) && board[row][col+1] != 0 && check(board, word, row, col+1, depth+1) {
			return true
		}
	}
	board[row][col] = tmp
	return false
}

// NOTE: this version looks shorter than but took more time to finish
// func check_2(board [][]byte, word string, row, col int, bMap [][]bool, depth int) bool {
//   if row < 0 || row == len(board) || col < 0 || col == len(board[0]) {
//     return false
//   }
//   if bMap[row][col] {
//     return false
//   }
//
//   if board[row][col] == word[depth] {
//     if depth == len(word)-1 {
//       return true
//     }
//     bMap[row][col] = true
//     if check(board, word, row-1, col, bMap, depth+1) ||
//       check(board, word, row+1, col, bMap, depth+1) ||
//       check(board, word, row, col-1, bMap, depth+1) ||
//       check(board, word, row, col+1, bMap, depth+1) {
//       return true
//     }
//   }
//   bMap[row][col] = false
//   return false
// }

func Run(board [][]byte, word string) bool {
	nRow, nCol := len(board), len(board[0])
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			c := check(board, word, i, j, 0)
			if c {
				return true
			}
		}
	}
	return false
}
