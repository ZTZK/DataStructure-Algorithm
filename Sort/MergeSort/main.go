package main

import "fmt"

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2 // Find the middle point of the array
	leftArray := arr[0:middle]
	rightArray := arr[middle:]
	return merge(mergeSort(leftArray), mergeSort(rightArray))
}

func merge(left []int, right []int) []int {
	mergeArray := []int{}             // Delare a slice to store the merged data
	leftPointer, rightPointer := 0, 0 // Declare two integers pointing to the first element of two slices

	for leftPointer < len(left) && rightPointer < len(right) {
		if left[leftPointer] > right[rightPointer] {
			mergeArray = append(mergeArray, right[rightPointer])
			rightPointer++
		} else {
			mergeArray = append(mergeArray, left[leftPointer])
			leftPointer++
		}
	}

	if leftPointer == len(left) { // If we have matched all the elements of the left slice
		mergeArray = append(mergeArray, right[rightPointer:]...)
	} else {
		mergeArray = append(mergeArray, left[leftPointer:]...)
	}

	return mergeArray
}

func main() {
	// This file shows the Merge Sort

	messyArray := []int{3, 2, 11, 33, 73, 56, 21, 13} // Declare a slice to store the unsorted numbers
	fmt.Println(mergeSort(messyArray))

}
