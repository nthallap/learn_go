package main

import "fmt"

type num []int

func main() {

	x := num{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range x {
		if i%2 == 0 {
			fmt.Printf("%d is even \n", i)
		} else {
			fmt.Printf("%d is odd \n", i)
		}

	}
}
