package hmap_test

import (
	hmap "data-structures/internal/hash-map"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	m := hmap.NewMap()

	for i := range 10 {
		m.Insert(fmt.Sprintf("test%d", i), i)
	}

	m.Print()
}
