// https://leetcode.com/problems/diameter-of-binary-tree/

package diameterofbinarytree

import "fmt"

func Input() (*TreeNode, int, int) {
	in_arr := []int{1, 2, 3, 0, 4, 0, 5}
	return NewTree(in_arr), 5, 4
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	ctn []*TreeNode
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func (this *Queue) EnQueue(v *TreeNode) {
	this.ctn = append(this.ctn, v)
}

func (this *Queue) isEmpty() bool {
	return len(this.ctn) == 0
}

func (this *Queue) DeQueue() *TreeNode {
	res := this.ctn[0]
	this.ctn = this.ctn[1:len(this.ctn)]
	return res
}
func NewQueue() *Queue {
	return &Queue{
		ctn: make([]*TreeNode, 0),
	}
}

func NewTree(input_array []int) *TreeNode {
	q := NewQueue()
	root := NewTreeNode(input_array[0])
	q.EnQueue(root)
	i := 1
	for !q.isEmpty() {
		n := q.DeQueue()
		if i < len(input_array) && input_array[i] != 0 {
			n.Left = NewTreeNode(input_array[i])
			q.EnQueue(n.Left)
		}
		if i < len(input_array) && input_array[i+1] != 0 {
			n.Right = NewTreeNode(input_array[i+1])
			q.EnQueue(n.Right)
		}
		i += 2
	}
	return root
}

func Bfs(queue *[]*TreeNode, x, y int) bool {
	new_queue := make([]*TreeNode, 0)
	has_x, has_y := make([]int, 0), make([]int, 0)
	not_empty := false
	// for i := range *queue {
	//   fmt.Printf("%d \t", (*queue)[i].Val)
	// }
	// fmt.Println()
	for i := range *queue {
		n := (*queue)[i]
		if n == nil {
			continue
		}
		if n.Val == x {
			has_x = append(has_x, i)
		}
		if (*queue)[i].Val == y {
			has_y = append(has_y, i)
		}
		new_queue = append(new_queue, n.Left, n.Right)
		if n.Left != nil || n.Right != nil {
			not_empty = true
		}
	}
	fmt.Println(has_x, has_y)
	for i := range has_x {
		for j := range has_y {
			if has_x[i]&1 == 0 && has_y[j] != has_x[i]+1 {
				return true
			}
			if has_x[i]&1 == 1 && has_y[j] != has_x[i]-1 {
				return true
			}
		}
	}
	if !not_empty {
		return false
	}
	return Bfs(&new_queue, x, y)
}

func Run(root *TreeNode, x, y int) bool {
	queue := []*TreeNode{root}
	return Bfs(&queue, x, y)
}
