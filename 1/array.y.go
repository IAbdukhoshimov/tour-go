package main

import "fmt"

func main() {
	var fnumber [6]float64
	fnumber[0] = 12
	fnumber[1] = 32
	fnumber[4] = 21
	fnumber[3] = 43
	fnumber[5] = 90

	for i, value := range fnumber {
		fmt.Println(value, i)
	}
}
