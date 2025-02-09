package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	elementToAdd  = 10
	indexToDelete = 1
)

func main() {
	originalSlice := make([]int, 10)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(50)
	}
	fmt.Println(originalSlice)
	fmt.Println(sliceExample(originalSlice))
	fmt.Println(addElements(originalSlice, elementToAdd))
	copyedSlice := copySlice(originalSlice)
	fmt.Println(copyedSlice)
	copyedSlice[1] = 50
	fmt.Printf("copyedSlice after changes: %v, originalSlice: %v\n", copyedSlice, originalSlice)

	fmt.Println(removeElements(originalSlice, indexToDelete))
}

func sliceExample(slice []int) []int {
	var newSlice []int
	for i := range slice {
		if slice[i]%2 == 0 {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}

func addElements(slice []int, num int) []int {
	slice = append(slice, num)
	return slice
}

func copySlice(slice []int) []int {
	copyedSlice := make([]int, len(slice))
	copy(copyedSlice, slice)
	return copyedSlice
}

func removeElements(slice []int, num int) []int {
	if num < 0 || num >= len(slice) {
		fmt.Printf("index %v out of range\n", num)
		return slice
	}
	return append(slice[:num], slice[num+1:]...)
}
