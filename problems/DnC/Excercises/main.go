package main

import (
	"fmt"
	"testing"
)

/*
Using Divide and Conquer
1. Add an Array of numbers
2. Find the largest number in an array
3. Do Binary Search from memory
*/

func addArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + addArray(arr[1:])
}

func numCompare(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestInArray(arr []int, high, low int) int {
	if high == low {
		return 0
	}
	if high == low+1 {
		return numCompare(high, low)
	}
	mid := len(arr) / 2
	leftMax := largestInArray(arr, low, mid)
	rightMax := largestInArray(arr, mid+1, high)

	return numCompare(leftMax, rightMax)
}

func LargestInArray(arr []int) int {
	low := 0
	high := len(arr) - 1
	return largestInArray(arr, low, high)
}

func TestLargestInArray(t *testing.T) {
	tests := map[string]struct {
		arr      []int
		expected int
	}{
		"test1": {arr: []int{1, 2, 3, 4, 5}, expected: 5},
		"test2": {arr: []int{10, 5, 8, 12, 3}, expected: 12},
		"test3": {arr: []int{-1, -5, -2, -8}, expected: -1},
		"test4": {arr: []int{0, 0, 0, 0}, expected: 0},
		"test5": {arr: []int{7}, expected: 7},
		"test6": {arr: []int{}, expected: 0},
		"test7": {arr: []int{100, 100, 99, 98}, expected: 100},
		"test8": {arr: []int{-10, -10, -10}, expected: -10},
	}

	for _, test := range tests {
		result :=LargestInArray(test.arr)
		if result != test.expected{
			t.Errorf("Case: %d Failed", )
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 23, 14, 65, 246, 74, 24}
	fmt.Println("Sum of Array using DnC =>", addArray(arr))
	fmt.Println("Max Val of Array using DnC =>", LargestInArray(arr))
}
