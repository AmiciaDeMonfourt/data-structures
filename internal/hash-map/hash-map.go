package hmap

import "fmt"

var (
	cap = 3 // table capacity
)

type hmap struct {
	schema []*List
	size   int64
}

func NewMap() *hmap {
	return &hmap{
		schema: make([]*List, cap),
		size:   0,
	}
}

func (m *hmap) Insert(key string, val interface{}) {
	hash := hash8bit(key)
	// if map is empty
	if m.schema[hash] == nil {
		m.schema[hash] = &List{}
	}
	m.schema[hash].Push(key, val)
}

func (m *hmap) Print() {
	for row := 0; row < cap; row++ {
		if m.schema[row] != nil {
			fmt.Println(row, ":", m.schema[row].Print())
			continue
		}
		fmt.Println(row, ":")
	}
}

func hash8bit(key string) (hash int) {
	for index, char := range key {
		hash += int(char) * (index)
	}
	return hash % cap
}
