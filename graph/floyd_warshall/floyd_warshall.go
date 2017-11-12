package floyd_warshall

import (
	"math"
)

func Search(graph [][]int) [][]int {
	dist := makeDist(graph)

	for t := 0; t < len(graph); t++ {
		for u := 0; u < len(graph); u++ {
			for v := 0; v < len(graph); v++ {
				newDist := dist[u][t] + dist[t][v]
				if newDist < dist[u][v] {
					dist[u][v] = newDist
				}
			}
		}
	}

	return dist
}

func makeDist(graph [][]int) [][]int {
	dist := make([][]int, len(graph))
	for i := 0; i < len(dist); i++ {
		dist[i] = make([]int, len(dist))
		for j := 0; j < len(dist[i]); j++ {
			if graph[i][j] != 0 {
				dist[i][j] = graph[i][j]
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
		dist[i][i] = 0
	}

	return dist
}
