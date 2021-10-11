package main

import "fmt"

func fact(number int) int {
	var result = 1
	for i := 1; i <= number; i++ {
		result = i * result
	}
	return result
	// fmt.Printf("%v this is the answer", result)
}

func main() {
	fmt.Println(fact(5))
	fmt.Println(fact(6))
	fmt.Println(fact(7))
}
