package hmap

import "fmt"

//2 8 256 512 1024 2048 5096
var (
	cap = 5096 // table capacity
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
	// generate a hashcode for the new node
	hash := hashcode(key)
	// if map is empty - initialize new list
	if m.schema[hash] == nil {
		m.schema[hash] = &List{}
	}

	m.schema[hash].Push(key, val)
}

func (m *hmap) Search(key string) interface{} {
	hash := hashcode(key)

	if m.schema[hash] == nil {
		panic(fmt.Sprintf("unknown record, key: %s", key))
	}

	return m.schema[hash].Search(key)
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

// simple func that generates bit hash code
// by key in the schema capacity range (0 to cap)
func hashcode(key string) (hash int) {
	for index, char := range key {
		hash += int(char) * (index)
	}
	return hash % cap
}
