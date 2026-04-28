package main

import "testing"

func TestPlotProblem(t *testing.T) {
	test := []struct {
		length   int
		width    int
		expected int
	}{
		{15, 10, 5},
		{10, 15, 5}, // commutative
		{100, 25, 25},
		{17, 13, 1},    // co-prime
		{12, 12, 12},   // equal
		{0, 10, 0},     // edge case: zero
		{0, 0, 0},      // edge case: both zero (will cause infinite recursion!)
		{-2, 3, -1},    // edge case: negative length
		{4, -5, -1},    // edge case: negative width
		{-83, -10, -1}, // edge case: negative length and width
	}
	for _, tt := range test {
		result := DividePlot(tt.length, tt.width)
		if result != tt.expected {
			t.Errorf("Divide Plot(%d, %d) = %d, want %d", tt.length, tt.width, result, tt.expected)
		}
	}
}

func BenchmarkDividePlot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DividePlot(123456, 3453)
	}
}
