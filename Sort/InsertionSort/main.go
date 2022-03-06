package main

import "fmt"

func main() {
	// This file will implement Insertion Sort

	messySlice := []int{13, 12, 14, 50, 33, 91, 3} // Declare a slice to store the unsorted numbers

	for i := 1; i < len(messySlice); i++ {
		preIndex := i - 1
		current := messySlice[i]
		for {
			if preIndex >= 0 && messySlice[preIndex] > current {
				messySlice[preIndex+1] = messySlice[preIndex]
				preIndex--
			} else {
				messySlice[preIndex+1] = current
				break
			}
		}

	}
	fmt.Println(messySlice)

}
