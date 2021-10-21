// https://leetcode.com/problems/insert-delete-getrandom-o1/

package insertdeleterandom

import "math/rand"

func Input() int {
	return 0
}

type RandomizedSet struct {
	index_map map[int]int
	keys      []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		index_map: make(map[int]int),
		keys:      make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.index_map[val]
	if ok {
		return false
	}
	this.keys = append(this.keys, val)
	this.index_map[val] = len(this.keys) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.index_map[val]
	if !ok {
		return false
	}
	last_v := this.keys[len(this.keys)-1]
	this.index_map[last_v] = idx
	this.keys[idx] = last_v
	delete(this.index_map, val)
	this.keys = this.keys[:len(this.keys)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.keys))
	return this.keys[idx]
}

func Run(t int) int {
	obj := Constructor()
	param_1 := obj.Insert(1)
	param_2 := obj.Remove(2)
	param_3 := obj.GetRandom()
	_ = param_1
	_ = param_2
	_ = param_3
	return t
}
