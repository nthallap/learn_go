package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var x = 10
	var y = 20.5
	var z = "naga"
	fmt.Println(x, y, z, runtime.GOOS)
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now())
	}
}
