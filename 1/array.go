package main

import "fmt"

func main() {
	var fnumber [5]float64

	fnumber[0] = 163
	fnumber[1] = 43
	fnumber[2] = 21
	fnumber[3] = 32
	fnumber[4] = 321
	fmt.Println(fnumber)
	fnumber2 := [5]float64{1, 2, 3, 2, 1}
	fmt.Println(fnumber2)
}
