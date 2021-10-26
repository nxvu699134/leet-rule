// https://leetcode.com/problems/count-complete-tree-nodes/

package countcompletetreenodes

func Input() *TreeNode {
	in_arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	return NewTree(in_arr)
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
		if i+1 < len(input_array) && input_array[i+1] != 0 {
			n.Right = NewTreeNode(input_array[i+1])
			q.EnQueue(n.Right)
		}
		i += 2
	}
	return root
}

func Run(root *TreeNode) *TreeNode {
	var Dfs func(*TreeNode)
	Dfs = func(root *TreeNode) {
		if root == nil || (root.Left == nil && root.Right == nil) {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		if root.Left != nil {
			Dfs(root.Left)
		}
		if root.Right != nil {
			Dfs(root.Right)
		}
	}
	Dfs(root)
	return root
}
