// https://leetcode.com/problems/min-stack/

package minstack

import "fmt"

func Input() int {
	return 0
}

//
// type MinStack struct {
//   cnt  []int
//   mins []int
// }
//
// func Constructor() MinStack {
//   return MinStack{
//     cnt:  make([]int, 0, 4),
//     mins: make([]int, 0, 4),
//   }
// }
//
// func (this *MinStack) Push(val int) {
//   this.cnt = append(this.cnt, val)
//   if len(this.mins) == 0 || val < this.mins[len(this.mins)-1] {
//     this.mins = append(this.mins, val)
//     return
//   }
//   this.mins = append(this.mins, this.mins[len(this.mins)-1])
// }
//
// func (this *MinStack) Pop() {
//   if len(this.cnt) == 0 {
//     return
//   }
//   this.cnt = this.cnt[:len(this.cnt)-1]
//   this.mins = this.mins[:len(this.mins)-1]
// }
//
// func (this *MinStack) Top() int {
//   if len(this.cnt) == 0 {
//     return 0
//   }
//   return this.cnt[len(this.cnt)-1]
// }
//
// func (this *MinStack) GetMin() int {
//   if len(this.cnt) == 0 {
//     return 0
//   }
//   return this.mins[len(this.mins)-1]
// }

type MinStack struct {
	cnt     []int
	cur_min int
}

func Constructor() MinStack {
	return MinStack{
		cnt:     make([]int, 0, 4),
		cur_min: (1 << 31) - 1,
	}
}

func (this *MinStack) Push(val int) {
	if val <= this.cur_min { // equal here cuz min can duplicate
		this.cnt = append(this.cnt, this.cur_min)
		this.cur_min = val
	}
	this.cnt = append(this.cnt, val)
}

func (this *MinStack) Pop() {
	if this.Top() == this.cur_min {
		this.cur_min = this.cnt[len(this.cnt)-2]
		this.cnt = this.cnt[:len(this.cnt)-2]
		return
	}
	this.cnt = this.cnt[:len(this.cnt)-1]
}

func (this *MinStack) Top() int {
	return this.cnt[len(this.cnt)-1]
}

func (this *MinStack) GetMin() int {
	return this.cur_min
}

func Run(_ int) int {
	minStack := Constructor()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	return 0
}
