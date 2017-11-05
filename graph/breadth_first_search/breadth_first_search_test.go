package breadth_first_search

import (
	"testing"
	"math"
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
	
	dist := Search(graph, 0)

	expected := []int{0,1,2,2,1,2,3,math.MaxInt32,math.MaxInt32}

	for i, _ := range dist {
		assert.Equal(expected[i], dist[i])
	}
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
