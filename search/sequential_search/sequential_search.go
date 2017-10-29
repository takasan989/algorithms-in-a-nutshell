package sequential_search

func Search(data []int, value int) bool {
	for _, v := range data {
		if v == value {
			return true
		}
	}
	return false
}