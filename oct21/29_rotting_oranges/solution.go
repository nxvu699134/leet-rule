// https://leetcode.com/problems/rotting-oranges/

package rottingoranges

func Input() [][]int {
	return [][]int{{0}}
	// return [][]int{
	//   {2, 1, 1},
	//   {0, 1, 1},
	//   {1, 0, 1},
	// }
}

type Queue struct {
	ctn []*Point
}

type Point struct {
	row, col int
}

func NewQueue() Queue {
	return Queue{
		ctn: make([]*Point, 0),
	}
}

func (this *Queue) EnQueue(v *Point) {
	this.ctn = append(this.ctn, v)
}

func (this *Queue) IsEmpty() bool {
	return len(this.ctn) == 0
}

func (this *Queue) Size() int {
	return len(this.ctn)
}

func (this *Queue) DeQueue() *Point {
	res := this.ctn[0]
	this.ctn = this.ctn[1:len(this.ctn)]
	return res
}

func Run(grid [][]int) int {
	q := NewQueue()
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 2 {
				q.EnQueue(&Point{i, j})
			}
		}
	}

	depth := 0
	for ; !q.IsEmpty(); depth++ {
		n := q.Size()
		for i := 0; i < n; i++ {
			p := q.DeQueue()
			if p.row > 0 && grid[p.row-1][p.col] == 1 {
				q.EnQueue(&Point{p.row - 1, p.col})
				grid[p.row-1][p.col] = 2
			}
			if p.row < len(grid)-1 && grid[p.row+1][p.col] == 1 {
				q.EnQueue(&Point{p.row + 1, p.col})
				grid[p.row+1][p.col] = 2
			}
			if p.col > 0 && grid[p.row][p.col-1] == 1 {
				q.EnQueue(&Point{p.row, p.col - 1})
				grid[p.row][p.col-1] = 2
			}
			if p.col < len(grid[0])-1 && grid[p.row][p.col+1] == 1 {
				q.EnQueue(&Point{p.row, p.col + 1})
				grid[p.row][p.col+1] = 2
			}
		}
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	if depth == 0 {
		return 0
	}
	return depth - 1
}
