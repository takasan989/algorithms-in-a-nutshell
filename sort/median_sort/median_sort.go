package median_sort

import (
	"../partition"
)

func Sort(data []int) {
	sort(data, 0, len(data) - 1)
}

func sort(data []int, left int, right int) {
	if right <= left {
		return
	}

	mid := (right - left) / 2;
	selectKth(data, left+mid, left, right)

	sort(data, left, left+mid-1)
	sort(data, left+mid+1, right)
}

func selectKth(data []int, k int, left int, right int) int {
	idx := left
	pivotIndex := partition.Partition(data, idx, left, right)

	if k == pivotIndex {
		return pivotIndex
	}

	if k < pivotIndex {
		return selectKth(data, k, left, pivotIndex - 1)
	} else {
		return selectKth(data, k, pivotIndex + 1, right)
	}
}