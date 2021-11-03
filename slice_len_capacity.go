package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(s, len(s), cap(s))

	s = s[:0]
	fmt.Println(s, len(s), cap(s))

	s = s[:3]
	fmt.Println(s, len(s), cap(s))

	s = s[:0]
	fmt.Println(s, len(s), cap(s))

	s = s[:3]
	fmt.Println(s, len(s), cap(s))
}
