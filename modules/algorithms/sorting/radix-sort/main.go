package radixsort

import "fmt"

// radixSort sorts an array using Radix Sort algorithm
func radixSort(arr []int) {
	if len(arr) == 0 {
		return // nothing to sort
	}

	// Step 1: Find the maximum number to know number of digits
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}

	// Step 2: Apply counting sort for each digit (exp = 1, 10, 100, ...)
	for exp := 1; maxVal/exp > 0; exp *= 10 {
		countingSortByDigit(arr, exp)
	}
}

// countingSortByDigit performs counting sort based on a specific digit (exp)
func countingSortByDigit(arr []int, exp int) {
	n := len(arr)

	// Output array to store sorted result for this digit
	output := make([]int, n)

	// Count array for digits 0–9
	count := make([]int, 10)

	// Step 1: Count occurrences of each digit
	for _, v := range arr {
		digit := (v / exp) % 10 // extract current digit
		count[digit]++
	}

	// Step 2: Convert count[] to cumulative count
	// This gives actual positions in output[]
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Step 3: Build output array (iterate backwards for stability)
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}

	// Step 4: Copy sorted result back into original array
	copy(arr, output)
}

func main() {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}

	radixSort(arr)

	fmt.Println(arr) // Output: [2 24 45 66 75 90 170 802]
}
