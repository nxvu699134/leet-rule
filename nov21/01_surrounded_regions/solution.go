// https://leetcode.com/problems/surrounded-regions/

package surroundedregions

import "fmt"

func Input() [][]byte {
	// return [][]byte{
	//   {'X', 'X', 'X', 'X'},
	//   {'X', 'O', 'O', 'X'},
	//   {'X', 'X', 'O', 'X'},
	//   {'X', 'O', 'X', 'X'},
	// }
	return [][]byte{
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	}
}

type Point struct {
	row, col int
}
type Queue struct {
	ctn []*Point
}

func (this *Queue) EnQueue(v *Point) {
	this.ctn = append(this.ctn, v)
}

func (this *Queue) IsEmpty() bool {
	return len(this.ctn) == 0
}

func (this *Queue) DeQueue() *Point {
	res := this.ctn[0]
	this.ctn = this.ctn[1:len(this.ctn)]
	return res
}
func NewQueue() *Queue {
	return &Queue{
		ctn: make([]*Point, 0),
	}
}

func solve(board [][]byte) {
	q := NewQueue()
	R := len(board)
	C := len(board[0])
	for i := 0; i < C; i++ {
		if board[0][i] == 'O' {
			q.EnQueue(&Point{0, i})
		}
	}
	if R > 0 {
		n := R - 1
		for i := 0; i < len(board[0]); i++ {
			if board[n][i] == 'O' {
				q.EnQueue(&Point{n, i})
			}
		}
	}

	for i := 0; i < R; i++ {
		if board[i][0] == 'O' {
			q.EnQueue(&Point{i, 0})
		}
	}
	if C > 0 {
		n := C - 1
		for i := 0; i < R; i++ {
			if board[i][n] == 'O' {
				q.EnQueue(&Point{i, n})
			}
		}
	}

	for !q.IsEmpty() {
		cur_p := q.DeQueue()
		if board[cur_p.row][cur_p.col] == 'o' {
			continue
		}
		board[cur_p.row][cur_p.col] = 'o'
		if cur_p.row > 0 && cur_p.row < R-1 {
			if cur_p.col < C-2 && board[cur_p.row][cur_p.col+1] == 'O' {
				q.EnQueue(&Point{cur_p.row, cur_p.col + 1})
			}
			if 1 < cur_p.col && board[cur_p.row][cur_p.col-1] == 'O' {
				q.EnQueue(&Point{cur_p.row, cur_p.col - 1})
			}
		}
		if cur_p.col > 0 && cur_p.col < C-1 {
			if cur_p.row < R-2 && board[cur_p.row+1][cur_p.col] == 'O' {
				q.EnQueue(&Point{cur_p.row + 1, cur_p.col})
			}
			if 1 < cur_p.row && board[cur_p.row-1][cur_p.col] == 'O' {
				q.EnQueue(&Point{cur_p.row - 1, cur_p.col})
			}
		}
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'o' {
				board[i][j] = 'O'
			}
		}
	}
}

func printBoard(board [][]byte) {
	fmt.Println()
	for i := range board {
		for j := range board[0] {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func Run(board [][]byte) int {
	solve(board)
	printBoard(board)
	return 0
}
