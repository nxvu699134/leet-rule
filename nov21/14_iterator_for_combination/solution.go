// https://leetcode.com/problems/iterator-for-combination/

package iteratorforcombination

import "fmt"

func Input() int {
	return 0
}

// Naive solution
// This solution beat 100% run time but only 25% memory used
// type Queue struct {
//   cnt []string
// }
//
// func NewQueue() *Queue {
//   return &Queue{
//     cnt: make([]string, 0),
//   }
// }
//
// func (this *Queue) EnQueue(val string) {
//   this.cnt = append(this.cnt, val)
// }
//
// func (this *Queue) DeQueue() string {
//   if len(this.cnt) == 0 {
//     return ""
//   }
//   ret := this.cnt[0]
//   this.cnt = this.cnt[1:]
//   return ret
// }
//
// func (this *Queue) IsEmpty() bool {
//   return len(this.cnt) == 0
// }
//
// type CombinationIterator struct {
//   comb *Queue
// }
//
// func generateCombinations(s []byte, length int, q *Queue, pre string) {
//   if len(s)+len(pre) < length {
//     return
//   }
//
//   if len(pre) == length {
//     q.EnQueue(pre)
//     return
//   }
//   for i := 0; i < len(s); i++ {
//     generateCombinations(s[i+1:], length, q, pre+string(s[i]))
//   }
// }
//
// func Constructor(characters string, combinationLength int) CombinationIterator {
//   q := NewQueue()
//   chars := []byte(characters)
//   generateCombinations(chars, combinationLength, q, "")
//   return CombinationIterator{
//     comb: q,
//   }
// }
//
// func (this *CombinationIterator) Next() string {
//   return this.comb.DeQueue()
// }
//
// func (this *CombinationIterator) HasNext() bool {
//   return !this.comb.IsEmpty()
// }
//
// func (this *CombinationIterator) Print() {
//   fmt.Println(*this.comb)
// }

// Bit manipulate solution
// This solution beat 100% run time but only 12.5% memory used
// (dont know why, I think this is better because I use only an integer to save pattern instead a string)
type Queue struct {
	cnt []int
}

func NewQueue() *Queue {
	return &Queue{
		cnt: make([]int, 0),
	}
}

func (this *Queue) EnQueue(val int) {
	this.cnt = append(this.cnt, val)
}

func (this *Queue) DeQueue() int {
	if len(this.cnt) == 0 {
		return 0
	}
	ret := this.cnt[0]
	this.cnt = this.cnt[1:]
	return ret
}

func (this *Queue) IsEmpty() bool {
	return len(this.cnt) == 0
}

func (this *Queue) Print() {
	for i := 0; i < len(this.cnt); i++ {
		fmt.Printf("%08b\n", this.cnt[i])
	}
}

type CombinationIterator struct {
	chars string
	comb  *Queue
}

func generateCombinations(numBits int, q *Queue, maskOffset int, pre int) {
	if numBits > maskOffset+1 { // because maskOffset contain zero "0"
		return
	}
	if numBits == 0 {
		q.EnQueue(pre)
		return
	}
	generateCombinations(numBits-1, q, maskOffset-1, pre|(1<<maskOffset))
	generateCombinations(numBits, q, maskOffset-1, pre)
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	q := NewQueue()
	generateCombinations(combinationLength, q, len(characters)-1, 0)
	return CombinationIterator{
		chars: characters,
		comb:  q,
	}
}

func (this *CombinationIterator) Next() string {
	ret := ""
	patt := this.comb.DeQueue()
	for mask, i := 1<<(len(this.chars)-1), 0; mask != 0; mask, i = mask>>1, i+1 {
		if mask&patt != 0 {
			ret += string(this.chars[i])
		}
	}
	return ret
}

func (this *CombinationIterator) HasNext() bool {
	return !this.comb.IsEmpty()
}

func (this *CombinationIterator) Print() {
	this.comb.Print()
}

func Run(n int) int {
	itr := Constructor("abcd", 2)
	// itr.Print()
	fmt.Println(itr.Next())
	fmt.Println(itr.Next())
	fmt.Println(itr.Next())
	return 0
}
