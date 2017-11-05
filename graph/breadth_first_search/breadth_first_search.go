package breadth_first_search

import (
	"math"
)

func Search(graph [][]int, s int) []int {
	queue := []int{}
	visited := make([]bool, len(graph))
	dist := make([]int, len(graph))
	for i, _ := range dist {
		dist[i] = math.MaxInt32
	}

	dist[s] = 0
	visited[s] = true
	queue = enqueue(queue, s)

	for len(queue) != 0 {
		var u int
		queue, u = dequeue(queue)

		for _, v := range neighbors(graph, u) {
			if !visited[v] {
				dist[v] = dist[u] + 1
				visited[v] = true
				queue = enqueue(queue, v)
			}
		}
	}
	return dist
}

func enqueue(queue []int, value int) []int {
	return append(queue, value)
}

func dequeue(queue []int) ([]int, int) {
	return queue[1:], queue[0]
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
