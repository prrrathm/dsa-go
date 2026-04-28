package sort



func InsertionSort(arr []int) []int {
	// Create a copy of the original slice
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	// Perform insertion sort on the copied slice
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

func InsertionSortMe(arr []int)[]int{
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	return sorted
}

func BubbleSort(arr []int) []int {
    n := len(arr)
    for i := range n-1 {
        swapped := false
        for j := range n-i-1 {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = true
            }
        }
        if !swapped {
            break
        }
    }
    return arr
}

// heapify maintains the heap property for a subtree rooted at index i
func heapify(arr []int, n int, i int) {
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
