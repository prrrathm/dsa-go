package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less []int
	var more []int
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
	result = append(quickSort(less), pivot)
	result = append(result, quickSort(more)...)
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// median (odd) = (n+1)/2 th element
	// median (even) = ((n/2) + (n+1)/2)/2
	var newArr []int
	if len(nums1) < 1 {
		newArr = nums2
	} else if len(nums2) < 1 {
		newArr = nums1
	} else {
		newArr = append(nums1, nums2...)
	}
	fmt.Println(newArr)
	if len(newArr) == 0 {
		return 0
	}
	if len(newArr) == 1 {
		return float64(newArr[0])
	}

	if len(newArr) == 2 {
		return float64((newArr[0] + newArr[1])) / 2
	}

	sorted := quickSort(newArr)
	fmt.Println(sorted)
	n := len(sorted)

	if n%2 == 0 {
		p1 := (n - 1) / 2
		p2 := (n-1)/2 + 1
		fmt.Println("p1, p2", sorted[p1], sorted[p2])
		return float64((sorted[p1] + sorted[p2])) / 2
	} else {
		p := (n - 1) / 2
		fmt.Println("p", sorted[p])
		return float64(sorted[p])
	}
}

func main() {
	arr1 := []int{2, 2, 4, 4}
	arr2 := []int{2, 2, 2, 4, 4}
	println("Answer", findMedianSortedArrays(arr1, arr2))
}
