package main

import "fmt"

type desc []string

func newDesk() desc {

	cards := desc{}
	suite := []string{"heart", "spade", "diamond"}
	vals := []string{"one", "two", "three", "king", "queen"}

	for _, val := range suite {
		for _, val2 := range vals {
			cards = append(cards, val2+" "+val)
		}
	}

	return cards
}

func (d desc) deal(num int) (desc, desc) {
	return d[:num], d[num:]
}

func (d desc) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}

func main() {
	cards := newDesk()
	// cards.print()
	front, back := cards.deal(3)
	front.print()
	back.print()
}
