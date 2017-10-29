package selection_sort

func Sort(data []int) {
	for right := len(data)-1; right > 0; right-- {
		index := maxIndex(data, 0, right)

		data[index], data[right] = data[right], data[index]
	}
}

func maxIndex(data []int, left int, right int) int {
	index := left

	for i := left+1; i <= right; i++ {
		if data[i] > data[index] {
			index = i
		}
	}
	return index
}