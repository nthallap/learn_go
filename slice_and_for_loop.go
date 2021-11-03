package main

import "fmt"

func main() {
	names := []string{}
	//{"naga", "chandra", "trijan", "navadeep"}
	names = append(names, "hello", "hwow", "are", "your")
	for i, name := range names {
		fmt.Println(i, name)
	}
}
