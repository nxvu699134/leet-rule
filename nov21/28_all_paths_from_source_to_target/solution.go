// https://leetcode.com/problems/all-paths-from-source-to-target/

package allpathsfromsourcetotarget

func Input() [][]int {
	// return [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	return [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
}

func visitGenerator(graph [][]int, result *[][]int) func(int, []int) {
	var visit func(int, []int)
	visit = func(nodeId int, prePath []int) {
		prePath = append(prePath, nodeId)
		if nodeId == len(graph)-1 {
			*result = append(*result, append([]int{}, prePath...))
			return
		}
		for _, nextNodeId := range graph[nodeId] {
			visit(nextNodeId, prePath)
		}
	}
	return visit
}

func allPathsSourceTarget(graph [][]int) [][]int {
	ret := make([][]int, 0)
	visitGenerator(graph, &ret)(0, []int{})
	// visit(0, []int{})
	return ret
}

func Run(g [][]int) [][]int {
	return allPathsSourceTarget(g)
}
