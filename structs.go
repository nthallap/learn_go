package main

import "fmt"

type Vertex struct {
	X int
	Y float64
}

type contactinfo struct {
	emailid string
	pin     int
}

type person struct {
	name    string
	contact contactinfo
}

func main() {
	v := Vertex{Y: 1}
	p := &v
	p.X = 25
	fmt.Println(p.X, p.Y)

	per := person{
		name: "nageswara Rao",
		contact: contactinfo{
			emailid: "naga@gmail.com",
			pin:     522000,
		},
	}
	fmt.Printf("%v", per)
}
