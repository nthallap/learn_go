package main

import "fmt"

type clolor string

func main() {
	myclr := clolor("red")
	fmt.Println("Hello world", myclr)
	fmt.Println(myclr.tellme("is beautiful"))

}

func (c clolor) tellme(data string) string {
	fmt.Println(c, "*****")
	return string(c) + " hello " + data
}
