// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

package findmininrotatedsortedarray

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

//
// func Run(root *TreeNode) int {
//   if root == nil {
//     return 0
//   }
//   cur_node := root
//   max_depth := 0
//   for cur_node.Left != nil {
//     max_depth++
//     cur_node = cur_node.Left
//   }
//   if max_depth == 0 {
//     return 1
//   }
//
//   count := 0
//   var Dfs func(*TreeNode, int) bool
//   Dfs = func(root *TreeNode, depth int) bool {
//     if depth < max_depth-1 {
//       if !Dfs(root.Right, depth+1) {
//         return Dfs(root.Left, depth+1)
//       }
//     }
//     if root.Right != nil {
//       return true
//     }
//     count++
//     if root.Left != nil {
//       return true
//     }
//     count++
//     return false
//   }
//   Dfs(root, 0)
//   return (1 << (max_depth + 1)) - 1 - count
// }

func Dfs(root *TreeNode, depth int, max_depth int, count *int) bool {
	if depth < max_depth-1 {
		if !Dfs(root.Right, depth+1, max_depth, count) {
			return Dfs(root.Left, depth+1, max_depth, count)
		}
	}
	if root.Right != nil {
		return true
	}
	*count++
	if root.Left != nil {
		return true
	}
	*count++
	return false
}

func Run(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur_node := root
	max_depth := 0
	for cur_node.Left != nil {
		max_depth++
		cur_node = cur_node.Left
	}
	if max_depth == 0 {
		return 1
	}
	count := 0
	Dfs(root, 0, max_depth, &count)
	return (1 << (max_depth + 1)) - 1 - count
}
