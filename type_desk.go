package main

import "fmt"

type Desk []string
type color string

func (d Desk) print() {

	for i, item := range d {
		fmt.Println(i, item)
	}
}

func (self Desk) myretun() Desk {
	return self
}
