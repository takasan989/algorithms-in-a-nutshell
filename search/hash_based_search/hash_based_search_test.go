package hash_based_search

import (
	"testing"
)

func TestSearch(t *testing.T) {
	data := []int{9, 24, 71, 31, 43, 32, 21, 64, 89, 57, 0, 78, 36, 61, 74, 97, 19, 6, 60, 76}
	table := loadTable(data)
	exists := Search(table, 57)

	if !exists {
		t.Errorf("Search #1 %v", data)
	}

	notExists := Search(table, 100)

	if notExists {
		t.Errorf("Search #2 %v", data)
	}
}

func TestLoadTable(t *testing.T) {
	data := []int{5, 2, 1, 4, 3, 0}
	table := loadTable(data)

	for _, list := range table {
		if len(list) != 1 {
			t.Errorf("loadTable list size %d", len(list))
		}
	}
}

func TestCreateTable(t *testing.T) {
	table := createTable()

	for _, list := range table {
		if len(list) != 0 {
			t.Errorf("CreateTable list size %d", len(list))
		}
	}
}

func TestHash(t *testing.T) {
	value := 123
	expected := 3

	actual := hash(value)

	if actual != expected {
		t.Errorf("hash #1 %d", actual)
	}
}
