package main

import "fmt"

func main() {
	numSlice := []int{4, 3, 2, 1, 1, 2}
	numSlice2 := make([]int, 5, 10)
	copy(numSlice2, numSlice)
	fmt.Println(numSlice)

}
