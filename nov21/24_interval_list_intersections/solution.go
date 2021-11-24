// https://leetcode.com/problems/interval-list-intersections/

package intervallistintersections

func Input() ([][]int, [][]int) {
	return [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
		[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ret := make([][]int, 0)
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		lower := max(firstList[i][0], secondList[j][0])
		upper := min(firstList[i][1], secondList[j][1])
		if upper-lower >= 0 {
			ret = append(ret, []int{lower, upper})
		}
		if firstList[i][1] > secondList[j][1] {
			j++
		} else if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			i++
			j++
		}
	}
	return ret
}

func Run(firstList [][]int, secondList [][]int) [][]int {
	return intervalIntersection(firstList, secondList)
}
