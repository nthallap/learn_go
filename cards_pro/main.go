package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

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

func (d desc) toString() string {
	return strings.Join([]string(d), ",")
}

func (d desc) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deckFromfile(filename string) desc {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("some thing went wrong")
	}

	arr := strings.Split(string(data), ",")
	return desc(arr)

}

func (d desc) ShuffleDesc() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := range d {
		randnum := r.Intn(len(d) - 1)
		d[i], d[randnum] = d[randnum], d[i]
	}
}

func main() {
	//cards := newDesk()
	// cards.print()
	// fmt.Printf("-----------------------")
	// front, back := cards.deal(3)
	// front.print()
	// fmt.Println("------------------------")
	// back.print()
	// fmt.Println("----------------------")
	// fmt.Println(back.toStrin())
	// cards.saveToFile("mycards")

	newcards := deckFromfile("mycards")
	newcards.print()
	newcards.ShuffleDesc()
	fmt.Println("-------------------")
	newcards.print()
}
