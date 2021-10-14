package main

import "fmt"

func main() {
	presAge := make(map[string]int)
	presAge["Islom Karimov"] = 42
	fmt.Println(len(presAge))
	fmt.Println(presAge)
}
