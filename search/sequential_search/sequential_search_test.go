package sequential_search

import (
	"testing"
)

func TestSearch(t *testing.T) {
	data := []int{9, 24, 71, 31, 43, 32, 21, 64, 89, 57, 0, 78, 36, 61, 74, 97, 19, 6, 60, 76}
	exists := Search(data, 57)

	if !exists {
		t.Errorf("search #1 %v", exists)
	}

	notExists := Search(data, 100)

	if notExists {
		t.Errorf("search #2 %v", notExists)
	}
}