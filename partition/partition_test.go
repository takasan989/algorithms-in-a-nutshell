package partition

import (
	"testing"
	"reflect"
)

func TestPartition(t *testing.T) {
	data := []int{5,2,4,8,3,7,9,1}
	expected := []int{2,1,3,8,5,7,9,4}

	Partition(data, 4, 0, len(data) - 1)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("partition 1 %v", data)
	}
}

func TestSwap(t *testing.T) {
	data := []int{5,2,3,1,4}
	expected := []int{5,2,4,1,3}

	swap(data, 2, len(data) - 1)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("swap right #1 %v", data)
	}
}