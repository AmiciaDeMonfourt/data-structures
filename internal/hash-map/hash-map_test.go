package hmap_test

import (
	hmap "data-structures/internal/hash-map"
	"fmt"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	m := hmap.NewMap()

	for i := range 10 {
		m.Insert(fmt.Sprintf("test%d", i), i)
	}

	m.Print()
}

func TestSearch(t *testing.T) {
	m := hmap.NewMap()

	for i := range 50_000_0 {
		m.Insert(fmt.Sprintf("test%d", i), i)
	}

	start := time.Now()
	fmt.Println("res: \n", m.Search("test56"))
	fmt.Println("search from 50_000_0 elements: ", time.Since(start).Microseconds(), "ms")
}
