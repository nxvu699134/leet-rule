// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

package guessnumber

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Input() []int {
	return []int{8, 5, 1, 7, 10, 12}
}

func NewNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func InsertNode(root *TreeNode, v int) {
	cur_node := root
	for {
		if v > cur_node.Val {
			if cur_node.Right == nil {
				cur_node.Right = NewNode(v)
				return
			}
			cur_node = cur_node.Right
		} else {

			if cur_node.Left == nil {
				cur_node.Left = NewNode(v)
				return
			}
			cur_node = cur_node.Left
		}
	}
}

func Run(preorder []int) *TreeNode {
	root := NewNode(preorder[0])
	for i := 1; i < len(preorder); i++ {
		InsertNode(root, preorder[i])
	}
	return root
}
