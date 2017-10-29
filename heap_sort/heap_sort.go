package heap_sort

func Sort(data []int) {
	buildHeap(data, len(data))

	for i := len(data)-1; i > 0; i-- {
		data[i], data[0] = data[0], data[i]
		heapify(data, 0, i)
	}
}

func buildHeap(data []int, len int) {
	for i := len/2; i >= 0; i-- {
		heapify(data, i, len)
	}
}

func heapify(data []int, root int, len int) {
	left := root * 2 + 1
	right := root * 2 + 2
	largest := root

	if left < len && data[left] > data[largest] {
		largest = left
	}

	if right < len && data[right] > data[largest] {
		largest = right
	}

	if largest != root {
		data[root], data[largest] = data[largest], data[root]
		heapify(data, largest, len)
	}
}