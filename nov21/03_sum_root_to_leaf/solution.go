// https://leetcode.com/problems/sum-root-to-leaf-numbers/

package sumroottoleaf

func Input() *TreeNode {
	in_arr := []int{1, 2, 3}
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

func Dfs(root *TreeNode, pre_sum int) int {
	pre_sum = pre_sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return pre_sum
	}
	ret := 0
	if root.Left != nil {
		ret += Dfs(root.Left, pre_sum)
	}

	if root.Right != nil {
		ret += Dfs(root.Right, pre_sum)
	}
	return ret
}

func sumNumbers(root *TreeNode) int {
	return Dfs(root, 0)
}

func Run(root *TreeNode) int {
	return sumNumbers(root)
}
