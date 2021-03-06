package median_sort

import (
	"testing"
	"reflect"
)

func TestSort(t *testing.T) {
	data := []int{1,4,8,9,11,15,7,12,13,6}
	expected := []int{1,4,6,7,8,9,11,12,13,15}

	Sort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Sort %v", data)
	}

	data = []int{9, 24, 71, 31, 43, 32, 21, 64, 89, 57, 0, 78, 36, 61, 74, 97, 19, 6, 60, 76}
	expected = []int{0, 6, 9, 19, 21, 24, 31, 32, 36, 43, 57, 60, 61, 64, 71, 74, 76, 78, 89, 97}

	Sort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Sort %v", data)
	}
}

func TestSelectKth(t *testing.T) {
	data := []int{5,3,1,2,4}
	expected := 3

	actual := selectKth(data, 2, 0, len(data) - 1)

	if data[actual] != expected {
		t.Errorf("select kth %v", actual)
	}
}
