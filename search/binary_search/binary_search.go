package binary_search

func Search(data []int, value int) bool {
	return search(data, value, 0, len(data)-1)
}

func search(data []int, value int, left int, right int) bool {
	if right < left {
		return false
	}

	mid := mid(left, right)

	if value == data[mid] {
		return true
	}

	if value < data[mid] {
		return search(data, value, left, mid-1)
	}

	return search(data, value, mid+1, right)
}

func mid(left int, right int) int {
	return left + (right-left)/2
}
