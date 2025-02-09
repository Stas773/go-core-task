package main

import "fmt"

func main() {
	slice1 := []int{65, 3, 3, 678, 64}
	slice2 := []int{64, 3, 4, 43}
	fmt.Println(sliceIntersection(slice1, slice2))
}

func sliceIntersection(slice1, slice2 []int) (bool, []int) {
	resultSlice := []int{}
	exists := false
	slice1Elements := make(map[int]bool)
	for _, v := range slice1 {
		slice1Elements[v] = true
	}
	for _, v := range slice2 {
		if slice1Elements[v] {
			resultSlice = append(resultSlice, v)
			delete(slice1Elements, v)
			exists = true
		}
	}
	return exists, resultSlice
}
