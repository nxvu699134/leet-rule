// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/

package minvaluetogetpositive

import (
	"sort"
)

func Input() [][]string {
	return [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}

	// return [][]string{
	//   {"David", "David0@m.co", "David1@m.co"},
	//   {"David", "David3@m.co", "David4@m.co"},
	//   {"David", "David4@m.co", "David5@m.co"},
	//   {"David", "David2@m.co", "David3@m.co"},
	//   {"David", "David1@m.co", "David2@m.co"},
	// }
}

func shouldMerge(accMap map[string]struct{}, arr []string) bool {
	for i := 1; i < len(arr); i++ {
		if _, ok := accMap[arr[i]]; ok {
			return true
		}
	}
	return false
}

func makeGraph(accounts [][]string) map[int][]int {
	relationMap := make(map[int][]int)
	for i := 0; i < len(accounts); i++ {
		m := make(map[string]struct{})
		for j := 1; j < len(accounts[i]); j++ {
			m[accounts[i][j]] = struct{}{}
		}

		relationMap[i] = make([]int, 0)
		for j := 0; j < len(accounts); j++ {
			if i == j || accounts[i][0] != accounts[j][0] {
				continue
			}
			if shouldMerge(m, accounts[j]) {
				relationMap[i] = append(relationMap[i], j)
			}
		}
	}
	return relationMap
}

func accountsMerge(accounts [][]string) [][]string {
	graph := makeGraph(accounts)
	ret := make([][]string, 0)
	for len(graph) != 0 {
		stack := make([]int, 0)
		mergeIdx := make([]int, 0)
		for k := range graph {
			stack = append(stack, k)
			break
		}

		for len(stack) != 0 {
			idx := stack[0]
			stack = stack[1:]
			if nextIdxes, ok := graph[idx]; ok {
				mergeIdx = append(mergeIdx, idx)
				stack = append(stack, nextIdxes...)
				delete(graph, idx)
			}
		}

		m := make(map[string]struct{})
		for i := 0; i < len(mergeIdx); i++ {
			for j := 1; j < len(accounts[mergeIdx[i]]); j++ {
				m[accounts[mergeIdx[i]][j]] = struct{}{}
			}
		}
		arr := make([]string, 0, len(m)+1)
		arr = append(arr, accounts[mergeIdx[0]][0])
		for mail := range m {
			arr = append(arr, mail)
		}
		sort.Strings(arr[1:])
		ret = append(ret, arr)
	}
	return ret
}

func Run(accs [][]string) [][]string {
	return accountsMerge(accs)
}
