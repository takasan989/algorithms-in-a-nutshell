package dijkstras_algorithm

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
	queue = enqueue(queue, s)

	for len(queue) != 0 {
		var v int
		queue, v = dequeue(queue)
		visited[v] = true
		
		for _, u := range neighbors(graph, v) {
			if visited[u] {
				continue
			}
			queue = enqueue(queue, u)
			updateDist(graph, dist, v, u)
		}
	}
	return dist
}

func updateDist(graph [][]int, dist []int, s int, v int) {
	weight := graph[s][v]

	if weight == 0 {
		return
	}

	newDist := dist[s] + weight

	if newDist < dist[v] {
		dist[v] = newDist
	}
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
