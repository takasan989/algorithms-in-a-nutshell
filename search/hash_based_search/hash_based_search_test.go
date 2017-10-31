package hash_based_search

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	tableSize := 6
	table := createTable(tableSize, []int{5, 2, 1, 4, 3})

	if len(table) != tableSize {
		t.Errorf("CreateTable size %d", len(table))
	}

	for _, list := range table {
		if len(list) != 0 {
			t.Errorf("CreateTable list size %d", len(list))
		}
	}
}

func TestHash(t *testing.T) {
	tableSize := 6
	value := 123
	expected := 3

	actual := hash(tableSize, value)

	if actual != expected {
		t.Errorf("hash #1 %d", actual)
	}
}
