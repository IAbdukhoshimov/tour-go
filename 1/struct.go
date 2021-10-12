package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 4}
	v.X = 56
	v.Y = 43
	fmt.Println(v.Y)
	fmt.Println(v.X)
}
