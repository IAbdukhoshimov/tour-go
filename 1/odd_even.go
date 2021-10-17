package main

import "fmt"

func main() {
	fmt.Println(odd_even(45))
}

func odd_even(num int) string {
	if num <= 1 {
		return "1"
	}

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			return "this is even"
		} else {
			return "this is odd "
		}
	}
	return "0"
}
