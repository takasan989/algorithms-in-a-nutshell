package depth_first_search

func Search(graph [][]int, s int) []bool {
	pred := make([]int, len(graph))
	visited := make([]bool, len(graph))

	visit(graph, visited, pred, s)

	for v, _ := range graph {
		if !visited[v] {
			visit(graph, visited, pred, v)
		}
	}
	return visited
}

func visit(graph [][]int, visited []bool, pred []int, v int) {
	visited[v] = true

	for _, u := range neighbors(graph, v) {
		if !visited[u] {
			pred[u] = v
			visit(graph, visited, pred, u)
		}
	}
}

func neighbors(graph [][]int, v int) []int {
	result := []int{}

	for i, _ := range graph[v] {
		if i != v && graph[v][i] != 0 {
			result = append(result, i)
		}
	}
	return result
}
