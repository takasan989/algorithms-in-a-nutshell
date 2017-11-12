package floyd_warshall

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	assert := assert.New(t)

	graph := [][]int{
		{0, 2, 0, 0, 4},
		{0, 0, 3, 0, 0},
		{0, 0, 0, 5, 1},
		{8, 0, 0, 0, 0},
		{0, 0, 0, 7, 0},
	}

	dist := Search(graph)

	expected := [][]int{
		{0, 2, 5, 10, 4},
		{16, 0, 3, 8, 4},
		{13, 15, 0, 5, 1},
		{8, 10, 13, 0, 12},
		{15, 17, 20, 7, 0},
	}

	assert.Equal(expected, dist)
}
