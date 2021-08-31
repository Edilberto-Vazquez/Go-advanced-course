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
			t.Errorf("Sum was incorret, got %d expected %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{2, 5, 5},
	}

	for _, v := range tables {
		max := GetMax(v.a, v.b)
		if max != v.n {
			t.Errorf("Getmax was incorrect, got %d, expected %d", max, v.n)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, v := range tables {
		fib := Fibonacci(v.a)
		if fib != v.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, v.n)
		}
	}
}
