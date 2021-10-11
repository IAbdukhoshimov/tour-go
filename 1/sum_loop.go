package main

import "fmt"

func sum(x int) int {
	result := 0
	for i := 0; i <= x; i++ {
		result += i
	}
	return result
}

func main() {
	fmt.Println(sum(100))
}
