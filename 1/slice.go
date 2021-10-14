package main

import "fmt"

func main() {
	numslice := []int{4, 3, 2, 3, 2, 34, 2, 3, 2, 1, 3, 21}
	numslice2 := numslice[3:8]
	fmt.Println("numslice2", numslice2)
	fmt.Println(numslice)
	for i, value := range numslice {
		fmt.Println(i, value)
	}
}
