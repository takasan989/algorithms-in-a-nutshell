package binary_search

import (
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	data := []int{9, 24, 71, 31, 43, 32, 21, 64, 89, 57, 0, 78, 36, 61, 74, 97, 19, 6, 60, 76}
	sort.Ints(data)
	exists := Search(data, 57)

	if !exists {
		t.Errorf("search #1 %v", exists)
	}

	notExists := Search(data, 100)

	if notExists {
		t.Errorf("search #2 %v", notExists)
	}
}

func TestMid(t *testing.T) {
	actual := mid(0, 9)
	expected := 4 //0,1,2,3,4,5,6,7,8,9

	if actual != expected {
		t.Errorf("mid #1 %d", actual)
	}

	actual = mid(0, 10)
	expected = 5

	if actual != expected {
		t.Errorf("mid #2 %d", actual)
	}

	actual = mid(1, 9)
	expected = 5

	if actual != expected {
		t.Errorf("mid #3 %d", actual)
	}
}
