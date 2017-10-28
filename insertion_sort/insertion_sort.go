package insertion_sort

import (
	"fmt"
)

func Sort(data[] int) {
	for i := 1; i < len(data); i++ {
		insert(data, i)
	}
}

func insert(data[]int, index int) error {
	idx, _ := insertionIndex(data, index)

	tmp := data[index]
	shift(data, idx, index)
	data[idx] = tmp
	return nil
}

func shift(data[]int, left int, right int) error {
	if len(data) == 0 {
		return fmt.Errorf("len(data) != 0")
	}

	for i := right; i > left; i-- {
		data[i] = data[i-1]
	}
	return nil
}

func insertionIndex(data []int, index int) (int, error) {
	if len(data) == 0 {
		return -1, fmt.Errorf("len(data) != 0")
	}

	for i := 0; i < index; i++ {
		if data[index] < data[i] {
			return i, nil
		}
	}
	return index, nil
}