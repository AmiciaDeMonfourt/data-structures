package hashMap

import "fmt"

type hMap struct {
	items []*Item
	size  uint
}

func NewMap() *hMap {
	return &hMap{
		items: make([]*Item, 0),
		size:  0,
	}
}

func (m *hMap) Add(key string, val interface{}) {
	index := m.size % m.hashCode(key)
	m.items[index] = NewItem(key, val)
	m.size++
}

func (m *hMap) Get(key string) *Item {
	index := m.size % m.hashCode(key)
	return m.items[index]
}

func (m *hMap) hashCode(key string) uint {
	var code int

	for c := range key {
		code += c
	}
	fmt.Println(key, "code", code)

	return uint(code)
}

func (m *hMap) Print() {
	for i := range m.size {
		fmt.Println(i, m.items[i])
	}
}
