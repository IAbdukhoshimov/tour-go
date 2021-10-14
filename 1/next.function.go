package main

import "fmt"

func main() {
	num1, num2 := nextValues(5)
	fmt.Println(num1, num2)
}

func nextValues(number int) (int, int) {
	return number + 1, number + 2
}
