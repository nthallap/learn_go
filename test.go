package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := 0
	y := 1
	var state int = 0
	return func() int {
		if state < 2 {
			if state == 0 {
				state++
				return 0
			}

			if state == 1 {
				state++
				return 1
			}
		}
		x, y = y, x+y

		return y
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println("done!")
}
