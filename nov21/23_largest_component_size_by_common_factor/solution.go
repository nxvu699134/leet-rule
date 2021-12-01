// https://leetcode.com/problems/largest-component-size-by-common-factor/

package largestcomponentsize

import (
	"math"
)

func Input() []int {
	return []int{2, 3, 6, 7, 4, 12, 21, 39}
	// return []int{20, 63, 50, 9}
}

func factorsOf(num int) map[int]struct{} {
	ret := make(map[int]struct{}, 0)
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			ret[i] = struct{}{}
			ret[num/i] = struct{}{}
		}
	}
	ret[num] = struct{}{}
	return ret
}

func hasToMerge(m1 map[int]struct{}, m2 map[int]struct{}) bool {
	for k := range m1 {
		if _, ok := m2[k]; ok {
			return true
		}
	}
	return false
}

func mergeMap(m1, m2 map[int]struct{}) map[int]struct{} {
	for k := range m1 {
		m2[k] = struct{}{}
	}
	return m2
}

func largestComponentSize(nums []int) int {
	factorMap := make(map[int]map[int]struct{})
	for i := 0; i < len(nums); i++ {
		factors := factorsOf(nums[i])
		for k := range factors {
			if _, ok := factorMap[k]; !ok {
				factorMap[k] = make(map[int]struct{})
			}
			factorMap[k][nums[i]] = struct{}{}
		}
	}
	groups := make([]map[int]struct{}, 0)
	for _, factors := range factorMap {
		newGroups := make([]map[int]struct{}, 0)
		for i := 0; i < len(groups); i++ {
			if hasToMerge(factors, groups[i]) {
				factors = mergeMap(factors, groups[i])
			} else {
				newGroups = append(newGroups, groups[i])
			}
		}
		groups = append(newGroups, factors)
	}
	max := len(groups[0])
	for i := 1; i < len(groups); i++ {
		if len(groups[i]) > max {
			max = len(groups[i])
		}
	}
	return max
}

func Run(nums []int) int {
	return largestComponentSize(nums)
}
