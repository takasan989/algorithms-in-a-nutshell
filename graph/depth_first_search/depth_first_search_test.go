package depth_first_search

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func graph() [][]int {
	graph := [][]int {
		{0,1,0,0,1,0,0,0,0},
		{1,0,1,0,0,0,0,0,0},
		{0,1,0,1,0,0,1,0,0},
		{0,0,1,0,1,0,0,0,0},
		{1,0,0,1,0,1,0,0,0},
		{0,0,0,0,1,0,0,0,0},
		{0,0,1,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,1},
		{0,0,0,0,0,0,0,1,0},
	}
	return graph
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	graph := graph()

	visited := Search(graph, 0)

	for _, r := range visited {
		assert.True(r)
	}
}

func TestVisit(t *testing.T) {
	assert := assert.New(t)
	graph := graph()

	visited := make([]bool, 9)
	pred := make([]int, 9)

	visit(graph, visited, pred, 8)

	assert.True(visited[7])
	assert.True(visited[8])
}

func TestNeighbors(t *testing.T) {
	assert := assert.New(t)

	graph := graph()

	neighbors := neighbors(graph, 2)

	assert.Len(neighbors, 3)
	assert.Contains(neighbors, 1)
	assert.Contains(neighbors, 3)
	assert.Contains(neighbors, 6)
}
