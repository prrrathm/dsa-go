/*
Given a plot of rectangular land that needs to be divided into squre plots of equal size find the maximum length of square that the plot can be divided into.
*/
package main

import (
	"fmt"
)

func DividePlot(length, width int) int {
	// Handling Edge Cases
	if length == 0 || width == 0 {
		return 0
	}
	if length < 0 || width < 0 {
		return -1
	}
	if length == width {
		return length
	}

	var smallerSide int
	var biggerSide int

	if length > width {
		smallerSide = width
		biggerSide = length
	} else {
		smallerSide = length
		biggerSide = width
	}

	return DividePlot(smallerSide, biggerSide-smallerSide)
}

func main() {
	var length int
	var width int

	fmt.Print("Enter Length of Plot => ")
	fmt.Scan(&length)

	fmt.Print("Enter Width of Plot => ")
	fmt.Scan(&width)

	fmt.Println("Maximum Size of Plot", DividePlot(length, width))
}
