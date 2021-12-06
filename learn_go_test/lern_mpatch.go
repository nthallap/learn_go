package main

import (
	"fmt"

	mpatch "C:\\Users\\nthall1x\\go\\src\\go-mpatch"
)

func methodA() int { return 1 }

//go:noinline
func methodB() int { return 2 }

func main() {

	patch, err := mpatch.PatchMethod(methodA, methodB)
	fmt.Println(methodA())
}
