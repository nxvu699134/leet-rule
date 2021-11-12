// https://leetcode.com/problems/remove-linked-list-elements/

package removelinkedlistelement

import "fmt"

func Input() (*ListNode, int) {
	return NewListNode([]int{6, 1, 2, 6, 3, 4, 5, 6}), 6
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(from []int) *ListNode {
	if len(from) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  from[0],
		Next: nil,
	}
	cur := head
	for i := 1; i < len(from); i++ {
		cur.Next = &ListNode{
			Val:  from[i],
			Next: nil,
		}
		cur = cur.Next
	}
	return head
}

func PrintList(head *ListNode) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}
	fmt.Println()
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val != val {
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next
		}
	}
	if head.Val == val {
		return head.Next
	}
	return head
}

func Run(head *ListNode, val int) int {
	PrintList(removeElements(head, val))
	return 0
}
