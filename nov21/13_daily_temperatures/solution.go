// https://leetcode.com/problems/daily-temperatures/

package dailytemperatures

import "fmt"

func Input() []int {
	return []int{73, 74, 75, 71, 69, 72, 76, 73}
}

type Stack struct {
	cnt []int
}

func NewStack() *Stack {
	return &Stack{
		cnt: make([]int, 0),
	}
}

func (this *Stack) Push(val int) {
	this.cnt = append(this.cnt, val)
}

func (this *Stack) Pop() int {
	if len(this.cnt) == 0 {
		return 0
	}
	ret := this.cnt[len(this.cnt)-1]
	this.cnt = this.cnt[:len(this.cnt)-1]
	return ret
}

func (this *Stack) Top() int {
	if len(this.cnt) == 0 {
		return 0
	}
	return this.cnt[len(this.cnt)-1]
}

func (this *Stack) IsEmpty() bool {
	return len(this.cnt) == 0
}

func (this *Stack) Print(t []int) {
	for i := range this.cnt {
		fmt.Printf("%d ", t[this.cnt[i]])
	}
	fmt.Println()
}

func dailyTemperatures(temperatures []int) []int {
	s := NewStack()
	ret := make([]int, len(temperatures))
	s.Push(0)
	for i := 1; i < len(temperatures); i++ {
		for !s.IsEmpty() && temperatures[i] > temperatures[s.Top()] {
			idx := s.Pop()
			ret[idx] = i - idx
		}
		s.Push(i)
	}
	return ret
}

func Run(temperatures []int) []int {
	return dailyTemperatures(temperatures)
}
