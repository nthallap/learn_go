package main

import "fmt"

func main() {
	var x = 10
	var p *int
	p = &x

	y := [...]string{"naga", "chandra", "navadeep", "trijan"}
	fmt.Println(x, p, *p, y)
}
