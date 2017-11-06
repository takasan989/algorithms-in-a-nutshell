package bellman_ford_algorithm

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
    assert := assert.New(t)
    
    graph := [][]int {
        {0,0,0,0,2},
        {0,0,0,-2,0},
        {0,-3,0,0,0},
        {0,0,6,0,0},
        {0,5,0,4,0},
    }
    
    dist := Search(graph, 0)
    
    assert.Equal(0, dist[0])
    assert.Equal(7, dist[1])
    assert.Equal(11, dist[2])
    assert.Equal(5, dist[3])
    assert.Equal(2, dist[4])
}