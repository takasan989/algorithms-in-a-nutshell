package hash_based_search

import (
	"algorithms-in-a-nutshell/search/sequential_search"
)

var tableSize = 6

func Search(table [][]int, value int) bool {
	h := hash(value)
	list := table[h]

	return sequential_search.Search(list, value)
}

func loadTable(data []int) [][]int {
	table := createTable()

	for _, v := range data {
		h := hash(v)
		table[h] = append(table[h], v)
	}
	return table
}

func createTable() [][]int {
	result := [][]int{}

	for i := 0; i < tableSize; i++ {
		list := []int{}
		result = append(result, list)
	}
	return result
}

func hash(value int) int {
	return value % tableSize
}
