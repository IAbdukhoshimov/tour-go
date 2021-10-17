package main

import "fmt"

func main() {
	fmt.Println(fib(10))
}

func fib(number int) int {
	if number <= 1 {
		return number
	}
	return fib(number-1) + fib(number-2)
}
