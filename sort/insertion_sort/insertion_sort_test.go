package insertion_sort

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
}

func TestInsert(t *testing.T) {
	data := []int{2,3,1,5,4}
	expected := []int{1,2,3,5,4}

	insert(data, 2)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("insert %v", data)
	}

	data = []int{1,4,9,6}
	expected = []int{1,4,6,9}

	insert(data, 1)
	insert(data, 2)
	insert(data, 3)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("insert %v", data)
	}

	data = []int{1,4,8,9,11,15,7,12,13,6}
	expected = []int{1,4,6,7,8,9,11,12,13,15}

	for i := 1; i < len(data); i++ {
		insert(data, i)
	}

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("insert %v", data)
	}
}

func TestInsertionIndexSuccess(t *testing.T) {
	expected := 2
	data := []int{1,4,8,9,11,15,7,12,13,6}
	idx, _ := insertionIndex(data, 6)

	if idx != expected {
		t.Errorf("insertion index %d is not %d", idx, expected)
	}

	data = []int{1,4,6,9}
	expected = 1

	idx, _ = insertionIndex(data, 1)

	if idx != expected {
		t.Errorf("insertion index %d is not %d", idx, expected)
	}
}

func TestInsertionIndexFailure(t *testing.T) {
	data := []int{}
	_, err := insertionIndex(data, 1)

	if err == nil {
		t.Errorf("insertion index error")
	}
}

func TestShift(t *testing.T) {
	expected := []int{1,2,2,3,4}
	data := []int{1,2,3,4,5}

	shift(data, 1, 4)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("shift %v", data)
	}

	expected = []int{1,4,8,8,9,11,15,12,13,6}
	data = []int{1,4,8,9,11,15,7,12,13,6}

	shift(data, 2, 6)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("shift %v", data)
	}
}