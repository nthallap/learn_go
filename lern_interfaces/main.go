package main

type bot interface {
	greetings() string
}

type englishBot struct{}
type teluguBot struct{}

func (englishBot) greetings() string {
	return "Hi, How are you ?"
}

func (teluguBot) greetings() string {
	return "Namaskaram, Eala unnaru ?"
}

func printGreeting(b bot) {
	print(b.greetings())
}
func main() {
	eb := englishBot{}
	tb := teluguBot{}
	printGreeting(eb)
	printGreeting(tb)
}
