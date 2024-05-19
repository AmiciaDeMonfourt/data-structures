package list_test

import (
	"data-structures/internal/list"
	"testing"
	"time"
)

func TestTiming(t *testing.T) {
	list := list.NewList()

	t.Log(list.GetSize())

	start := time.Now()

	for i := 1; i < 10_000_0; i++ {
		list.PushAfterNode("test", uint(i-24))

	}

	elapsed := time.Since(start)

	t.Log(elapsed)
}
