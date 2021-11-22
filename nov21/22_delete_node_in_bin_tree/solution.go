// https://leetcode.com/problems/delete-node-in-a-bst/

package deletenodeinbintree

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

//
// func deleteNode(root *TreeNode, key int) *TreeNode {
//   if root == nil {
//     return nil
//   }
//
//   if key < root.Val {
//     root.Left = deleteNode(root.Left, key)
//   } else if key > root.Val {
//     root.Right = deleteNode(root.Right, key)
//   } else {
//     if root.Left == nil && root.Right == nil {
//       return nil
//     }
//     if root.Left == nil && root.Right != nil {
//       return root.Right
//     }
//     if root.Left != nil && root.Right == nil {
//       return root.Left
//     }
//
//     // find the greatest in sub left
//     greatestNode := root.Left
//     for greatestNode.Right != nil {
//       greatestNode = greatestNode.Right
//     }
//     root.Val = greatestNode.Val
//     root.Left = deleteNode(root.Left, root.Left.Val)
//   }
//   return root
// }

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// find the greatest in sub left
		greatestNode := root.Left
		for greatestNode.Right != nil {
			greatestNode = greatestNode.Right
		}
		// This solution is run faster but That will make the tree unbalance
		greatestNode.Right = root.Right
		return root.Left
	}
	return root
}
func Run(root *TreeNode, key int) int {
	deleteNode(root, key)
	return 0
}
