package hash_based_search

func createTable(size int, data []int) [][]int {
	result := [][]int{}

	for i := 0; i < size; i++ {
		list := []int{}
		result = append(result, list)
	}
	return result
}

func hash(tableSize int, value int) int {
	return value % tableSize
}
