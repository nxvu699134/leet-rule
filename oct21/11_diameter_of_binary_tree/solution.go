// https://leetcode.com/problems/diameter-of-binary-tree/

package diameterofbinarytree

func Input() *TreeNode {
	in_arr := []int{1, 2, 3, 4, 5}
	return NewTree(&in_arr)
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(input_array *[]int) *TreeNode {
	return &TreeNode{}
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 1
	}
	left := dfs(root.Left, res)
	right := dfs(root.Right, res)
	t := left + right
	if t > *res {
		*res = t
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func Run(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}
