package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	// total := Sum(5, 5)

	// if total != 11 {
	// 	t.Errorf("sum was incorrect, got %d expected %d", total, 11)
	// }

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{5, 5, 10},
		{5, 8, 13},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("sum was incorret, got %d expected %d", total, item.n)
		}
	}
}
