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
	contact contactinfo // contactinfo contactinfo or simply contactinfo
}

type person2 struct {
	name        string
	contactinfo // contactinfo contactinfo or simply contactinfo
}

func (p person2) print() {

	fmt.Printf("\n %+v", p)

}

func (p *person2) updatename(fname string) {
	(*p).name = fname
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

	per2 := person2{
		name: "nageswara Rao",
		contactinfo: contactinfo{
			emailid: "naga@gmail.com",
			pin:     522000,
		},
	}
	perpinter := &per2
	perpinter.updatename("chandra sekhar")
	per2.print()
	// we can call updatename with per2 type variable directly with out creating a pointer var  type

	per2.updatename("navadeep")
	per2.print()
}
