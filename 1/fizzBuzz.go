package main

import (
	"fmt"
)

func fizzBuzz(x int) string {
	if x%15 == 0 {
		return "fizzBuzz"
	} else if x%3 == 0 {
		return "fizz"
	} else if x%15 == 0 {
		return "Buzz"
	}
	return "it"
}
func main() {
	fmt.Println(fizzBuzz(15))
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(45))
}
