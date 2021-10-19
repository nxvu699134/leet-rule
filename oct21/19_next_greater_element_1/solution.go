// https://leetcode.com/problems/next-greater-element-i/

package nextgreaterelement

import "fmt"

func Input() ([]int, []int) {
	return []int{4, 1, 2}, []int{1, 3, 4, 2}
}

func Run(num1 []int, num2 []int) []int {
	m := make(map[int]int)
	for i := range num2 {
		m[num2[i]] = i
	}
	for k := range m {
		fmt.Println(k)
	}
	res := make([]int, 0)
	for i := range num1 {
		idx := m[num1[i]]
		c := -1
		for j := idx + 1; j < len(num2); j++ {
			if num2[j] > num2[idx] {
				c = num2[j]
				break
			}
		}
		res = append(res, c)
	}
	return res
}
