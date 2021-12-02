// https://leetcode.com/problems/odd-even-linked-list/

package oddevenlinkedlist

import "fmt"

func Input() *ListNode {
	n := 10
	root := &ListNode{0, nil}
	cur := root
	for i := 1; i < n; i++ {
		cur.Next = &ListNode{i, nil}
		cur = cur.Next
	}
	return root
}

func Print(head *ListNode) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	oddHead := head.Next
	head.Next = nil

	cur := make(map[bool]*ListNode)
	cur[false] = head
	cur[true] = oddHead
	isOdd := false
	for cur[false].Next != cur[true].Next {
		cur[false].Next, cur[true].Next = cur[true].Next, cur[false].Next
		cur[isOdd] = cur[isOdd].Next
		isOdd = !isOdd
	}
	cur[false].Next = oddHead
	return head
}

func Run(head *ListNode) int {
	Print(head)
	h := oddEvenList(head)
	Print(h)
	return 0
}
