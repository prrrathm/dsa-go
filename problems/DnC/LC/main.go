package main

import (
	"fmt"
	"strconv"
)

// func charFrequency(s string) map[rune]int {
// 	freq := make(map[rune]int)
// 	for _, char := range s {
// 			freq[char]++
// 	}
// 	return freq
// }

// func maxDifference(s string) int {
// 	freq := charFrequency(s) 
// 	fmt.Println(freq)

// 	largestOdd := 0
// 	smallestEven := 0
// 	foundEvenFirst := false

// 	for _, item := range freq{
// 		if item % 2 == 0{
// 			if !foundEvenFirst {
// 				smallestEven = item
// 				foundEvenFirst = true
// 			} else if item < smallestEven{
// 				smallestEven = item
// 			}
// 		} else {
// 			if item > largestOdd{
// 				largestOdd = item
// 			} 
// 		}
// 	}
// 	// print(smallestEven, largestOdd)
// 	return largestOdd - smallestEven
// }

func reverseBitsStr(numStr string) string{
	if len(numStr) < 1{
		return numStr
	}
	return reverseBitsStr(numStr[1:]) + string(numStr[0])
}

func reverseBits(num uint32) uint32 {
	binaryStr := fmt.Sprintf("%032s", strconv.FormatUint(uint64(num), 2))
	reverseBinaryStr := reverseBitsStr(binaryStr)

	val, _ := strconv.ParseUint(reverseBinaryStr, 2, 32)

	return uint32(val)
}

func main(){
	// s := "aaaaabbc"
	// fmt.Println("Max Difference", maxDifference(s))
	binary := "00000010100101000001111010011100"
	val, err := strconv.ParseUint(binary, 2, 32)
	if err != nil {
			// handle error
	}
	n := uint32(val)

	fmt.Print(reverseBits(n))
}