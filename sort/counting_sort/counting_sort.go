package counting_sort

func Sort(data []int) {
	len := max(data)+1
	bucket := make([]int, len)

	for i, _ := range data {
		bucket[data[i]]++
	}

	index := 0
	for i := 0; i < len; i++ {
		for bucket[i] > 0 {
			data[index] = i
			index++
			bucket[i]--
		}
	}
}

func max(data []int) int {
	max := -1
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}