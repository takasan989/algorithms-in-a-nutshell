package quick_sort

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

	pivot := partition.Partition(data, left, left, right)

	sort(data, left, pivot - 1)
	sort(data, pivot + 1, right)
}