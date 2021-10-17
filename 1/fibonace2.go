package main

import "fmt"

func main() {
	fmt.Println(fib(10))
}

func fib(number int) int {
	if number <= 1 {
		return number
	}
	var result int = 0
	var first int = 1
	var second int = 0
	for i := 1; i <= number; i++ {
		result = first + second
		first = second
		second = result
	}

	return result

}
