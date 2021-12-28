package main

import (
	"fmt"
	"os"
)

func main() {

	val, err := os.Open("naga.rar")
	fmt.Println(val, err)

}
