package fib

import "testing"

func TestFib(t *testing.T) {
	var fibTests = []struct {
		n        int
		expected int
	}{{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13}}
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func TestFib2(t *testing.T) {

	testTable := map[string]struct {
		n        int
		expected int
	}{
		"success1": {1, 1},
		"success2": {2, 1},
		"error":    {1, 3},
	}

	for name, tt := range testTable {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("name: %s, Fib(%d): expected %d, actual: %d", name, tt.n, tt.expected, actual)
		}
	}
}

var result int

func BenchmarkFib(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Fib(n)
	}
	result = r
}
