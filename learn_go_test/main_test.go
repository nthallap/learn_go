package main

import (
	"fmt"
	"testing"
)

type Input struct {
	x        int
	y        int
	expected int
}

var results = []Input{
	{2, 3, 5},
	{5, 7, 12},
	{7, -2, 5},
}

// func TestAdd(t *testing.T) {
// 	res := Add(2, 3)
// 	if res != 5 {
// 		t.Error("addition is wrong")
// 	}
// }

func TestAdd(t *testing.T) {
	for _, tc := range results {
		var naga = 20
		fmt.Println("testing", naga)
		val := Add(tc.x, tc.y)

		fmt.Printf("x = %d, y = %d, expected = %d, output=%d \n ", tc.x, tc.y, tc.expected, val)
		if val != tc.expected {
			t.Error("wrong output")
		}
	}
}
