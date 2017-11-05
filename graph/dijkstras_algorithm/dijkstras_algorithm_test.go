package dijkstras_algorithm

import (
	"testing"
	"math"
	"github.com/stretchr/testify/assert"
)

func graph() [][]int {
	graph := [][]int {
		{0,2,0,0,4},
		{0,0,3,0,0},
		{0,0,0,5,1},
		{8,0,0,0,0},
		{0,0,0,7,0},
	}
	return graph
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)

	graph := graph()

	dist := Search(graph, 0)

	assert.Equal(0, dist[0])
	assert.Equal(2, dist[1])
	assert.Equal(5, dist[2])
	assert.Equal(10, dist[3])
	assert.Equal(4, dist[4])
}

func TestUpdateDist(t *testing.T) {
	assert := assert.New(t)

	graph := graph()
	dist := []int{0, math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32}

	updateDist(graph, dist, 0, 1)

	assert.Equal(2, dist[1])

	updateDist(graph, dist, 0, 3)

	assert.Equal(math.MaxInt32, dist[3])
}
