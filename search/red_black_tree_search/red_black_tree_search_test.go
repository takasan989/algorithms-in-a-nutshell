package red_black_tree_search

import (
    "testing"
)

func TestSearch(t *testing.T) {
 	data := []int{9, 24, 71, 31, 43, 32, 21, 64, 89, 57, 0, 78, 36, 61, 74, 97, 19, 6, 60, 76}
 	tree := NewTree(data)
	exists := Search(tree, 57)

	if !exists {
		t.Errorf("Search #1 %v", data)
	}

	notExists := Search(tree, 100)

	if notExists {
		t.Errorf("Search #2 %v", data)
	}
}