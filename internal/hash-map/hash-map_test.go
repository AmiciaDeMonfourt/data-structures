package hmap_test

import (
	hmap "data-structures/internal/hash-map"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHashMap_Insert(t *testing.T) {
	m := hmap.NewMap()

	for i := range 10 {
		m.Insert(fmt.Sprintf("test%d", i), i)
	}

	m.Print()
}

func TestHashMap_Remove(t *testing.T) {
	m := hmap.NewMap()

	for i := range 10 {
		m.Insert(fmt.Sprintf("test%d", i), i)
	}

	assert.Equal(t, 5, m.Search("test5"))

	m.Delete("test5")

	assert.Panics(t, func() {
		m.Search("test5")
	})

	// m.Print()
}

func TestHashMap_Search(t *testing.T) {
	m := hmap.NewMap()

	for i := range 10 {
		m.Insert(fmt.Sprintf("test%d", i), i)
	}

	assert.Equal(t, 5, m.Search("test5"))

	assert.Panics(t, func() {
		m.Search("panic")
	})
}

func TestHashMap_SearchRate(t *testing.T) {
	m := hmap.NewMap()

	for i := range 10_000_0 {
		m.Insert(fmt.Sprintf("test%d", i), i)
	}

	start := time.Now()
	fmt.Println("res: \n", m.Search("test34534"))
	fmt.Println("search from 50_000_0 elements: ", time.Since(start).Microseconds(), "ms. Capacity 256")
}
