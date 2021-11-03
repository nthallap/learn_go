package main

import "fmt"

var m = map[string]int{
	"naga": 123,
}

func main() {
	k := make(map[int]int)

	k[5] = 5
	k[10] = 10
	fmt.Println(m, k)
}
