package bellman_ford_algorithm

import (
	"math"
)

func Search(graph [][]int, s int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt32
	}

	dist[s] = 0
	edges := edges(graph)

	for i := 0; i < n; i++ {
		if i == n {
			return nil
		}
		changed := false

		for _, edge := range edges {
			u := edge[0]
			v := edge[1]

			newDist := dist[u] + graph[u][v]
			if newDist < dist[v] {
				dist[v] = newDist
				changed = true
			}
		}

		if !changed {
			break
		}
	}
	return dist
}

func edges(graph [][]int) [][]int {
	result := [][]int{}

	for u := 0; u < len(graph); u++ {
		for v := 0; v < len(graph[u]); v++ {
			if graph[u][v] != 0 {
				result = append(result, []int{u, v})
			}
		}
	}
	return result
}
