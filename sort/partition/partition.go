package partition

func Partition(data []int, pivot int, left int, right int) int {
	pivotValue := data[pivot]
	swap(data, pivot, right)

	target := left

	for i := left; i < right; i++ {
		if data[i] < pivotValue {
			data[i], data[target] = data[target], data[i]
			target++
		}
	}

	swap(data, target, right)
	return target
}

func swap(data []int, index int, right int) {
	data[index], data[right] = data[right], data[index]
}