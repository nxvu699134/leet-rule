// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

package inorderpostordertotree

func Input() ([]int, []int) {
	return []int{4, 8, 2, 5, 1, 6, 3, 7},
		[]int{8, 4, 5, 2, 6, 7, 3, 1}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := NewTreeNode(postorder[len(postorder)-1])
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}

func Run(inorder, postorder []int) int {
	buildTree(inorder, postorder)
	return 0
}
