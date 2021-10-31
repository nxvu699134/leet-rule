// https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/

package flattendoublylinkedlist

import "fmt"

func Input() *Node {
	// n1, n2, n3, n4, n5, n6, n7, n8, n9 := NewNode(1), NewNode(2), NewNode(3), NewNode(4), NewNode(5), NewNode(6), NewNode(7), NewNode(8), NewNode(9)
	// n1.Next = n2
	// n2.Prev = n1
	//
	// n2.Next = n3
	// n3.Prev = n2
	//
	// n3.Next = n4
	// n4.Prev = n3
	//
	// n3.Child = n5
	//
	// n5.Next = n6
	// n6.Prev = n5
	//
	// n6.Next = n7
	// n7.Prev = n6
	//
	// n7.Next = n8
	// n8.Prev = n7
	//
	// n7.Child = n9
	//
	// return n1

	n1, n2, n3, n4 := NewNode(1), NewNode(2), NewNode(3), NewNode(4)
	n1.Child = n2
	n2.Child = n3
	n3.Next = n4
	n4.Prev = n3
	return n1

	// n1, n2, n3, n4, n5 := NewNode(1), NewNode(2), NewNode(3), NewNode(4), NewNode(5)
	// n1.Child = n2
	//
	// n2.Child = n5
	// n2.Next = n3
	// n3.Prev = n2
	//
	// n3.Next = n4
	// n4.Prev = n3
	// return n1
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

type Queue struct {
	ctn []*Node
}

func NewQueue() Queue {
	return Queue{
		ctn: make([]*Node, 0),
	}
}

func (this *Queue) EnQueue(v *Node) {
	this.ctn = append(this.ctn, v)
}

func (this *Queue) DeQueue() *Node {
	res := this.ctn[0]
	this.ctn = this.ctn[1:len(this.ctn)]
	return res
}

func (this *Queue) IsEmpty() bool {
	return len(this.ctn) == 0
}

func NewNode(val int) *Node {
	return &Node{
		Val:   val,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
}

func flatten(root *Node) *Node {
	q := NewQueue()
	for cur_node := root; cur_node != nil; cur_node = cur_node.Next {
		if cur_node.Child != nil {
			q.EnQueue(cur_node)
		}
	}

	for !q.IsEmpty() {
		n := q.DeQueue()
		fmt.Println(n)
		last_node := n.Child
		for ; ; last_node = last_node.Next {
			if last_node.Child != nil {
				q.EnQueue(last_node)
			}
			if last_node.Next == nil {
				break
			}
		}
		//at tail
		last_node.Next = n.Next
		if n.Next != nil {
			n.Next.Prev = last_node
		}
		//at head
		n.Next = n.Child
		n.Child.Prev = n
		n.Child = nil
	}
	return root
}

func Run(root *Node) int {
	for cur := root; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}

	r := flatten(root)
	cur := r

	for ; ; cur = cur.Next {
		fmt.Printf("%d\t", cur.Val)
		if cur.Next == nil {
			break
		}
	}

	fmt.Println()

	for ; ; cur = cur.Prev {
		fmt.Printf("%d\t", cur.Val)
		if cur.Prev == nil {
			break
		}
	}
	fmt.Println()
	return 0
}
