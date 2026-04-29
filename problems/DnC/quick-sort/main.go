package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := []int{}
	more := []int{}

	for _, i := range arr[1:] {
		if i <= pivot {
			less = append(less, i)
		}
	}

	for _, i := range arr[1:] {
		if i > pivot {
			more = append(more, i)
		}
	}
	var result []int
	result = append(QuickSort(less), pivot)
	result = append(result, QuickSort(more)...)
	return result
}

func main(){
	arr := []int{1,2,4,23,67}
	fmt.Printf("This is the sorted array => %d", arr)
}