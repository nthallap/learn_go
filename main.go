package main

import "fmt"

func main() {
	names := Desk{"naga", "rama"}

	names.print()
	fmt.Println(names.myretun())
	// for i, item := range names {
	// 	fmt.Println(i, item)
	// }
}
