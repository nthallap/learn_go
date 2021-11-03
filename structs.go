package main

import "fmt"

type Vertex struct {
	X int
	Y float64
}

func main() {
	v := Vertex{Y: 1}
	p := &v
	p.X = 25
	fmt.Println(p.X, p.Y)
}
