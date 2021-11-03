package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var name string
	var age int
	fmt.Println("Enter name and age: ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("My name is %s and my age is %d", name, age)
	fmt.Println("--------------------------")
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("Enter user name second time")
	i, j := buf.ReadString('\n')
	fmt.Printf("my values are %s, %s", i, j)
}
