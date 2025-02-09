package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(sliceDifference(slice1, slice2))
}

func sliceDifference(slice1, slice2 []string) []string {
	resultSlice := []string{}
	slice2Elements := make(map[string]bool)
	for _, v := range slice2 {
		slice2Elements[v] = true
	}
	for _, v := range slice1 {
		if !slice2Elements[v] {
			resultSlice = append(resultSlice, v)
		}
	}
	return resultSlice
}
