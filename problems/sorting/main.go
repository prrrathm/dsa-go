package main

/*
Major Sorting Algorithms
1. Selection Sort
2. Insertion Sort
3. Bubble Sort
4. Merge Sort
5. Heapsort
6. Quicksort
*/

import (
	"fmt"
	// "dsa-go/modules/algorithms"
)

func SelectionSort(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	for i := range sorted {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j] < sorted[i] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	return sorted
}

func InsertionSort(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j] > key {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}
	return sorted
}

func BubbleSort(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	for i := range len(sorted) - 1 {
		swapped := false
		for j := range len(sorted) - i - 1 {
			if sorted[j] > sorted[j+1] {
				sorted[j+1], sorted[j] = sorted[j], sorted[j+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return sorted
}

// heapify maintains the heap property for a subtree rooted at index i
func heapify(arr []int, n int, i int) {
	// fmt.Println("Largest in recursion =>", i, arr[i])
	largest := i       // Initialize largest as root
	left := 2*i + 1    // left child
	right := 2*i + 2   // right child

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}

// HeapSort implements the heap sort algorithm
func HeapSort(arr []int) []int {
	// n := len(arr)
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	// Build max heap
	for i := len(sorted)/2 - 1; i >= 0; i-- {
	// fmt.Println("Largest in loop =>", i, sorted[i])
		heapify(sorted, len(sorted), i)
	}

	// Extract elements from heap one by one
	for i := len(sorted) - 1; i >= 0; i-- {
		// Move current root to end
		sorted[0], sorted[i] = sorted[i], sorted[0]
		// call max heapify on the reduced heap
		heapify(sorted, i, 0)
	}
	return sorted
}


func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	arr2 := []int{64, 34, 25, 12, 22, 11, 90, 10, 3, 2}
	fmt.Println("Original array:", arr)
	fmt.Println("Insertion sort:", InsertionSort(arr))
	fmt.Println("Selection sort:", SelectionSort(arr))
	fmt.Println("Bubble sort:", BubbleSort(arr))
	fmt.Println("Bubble sort:", HeapSort(arr2))
}
