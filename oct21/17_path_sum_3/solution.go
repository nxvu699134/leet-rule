// https://leetcode.com/problems/diameter-of-binary-tree/

package diameterofbinarytree

func Input() (*TreeNode, int) {
	in_arr := []int{1, 0, 2, 0, 3, 0, 4, 0, 5}
	return NewTree(in_arr), 3
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

func dfs(root *TreeNode, accum, target int, is_start bool) int {
	if root == nil {
		return 0
	}

	new_accum := accum + root.Val
	c := 0
	if new_accum == target {
		c = 1
	}
	res := c + dfs(root.Left, new_accum, target, false) + dfs(root.Right, new_accum, target, false)
	if is_start {
		res += dfs(root.Left, 0, target, true) + dfs(root.Right, 0, target, true)
	}
	// fmt.Printf("%d %d %d %d\n", res, accum, new_accum, root.Val)
	return res
}

func Run(root *TreeNode, target int) int {
	return dfs(root, 0, target, true)
}
