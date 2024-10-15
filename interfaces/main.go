package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating English greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating Spanish greeting
	return "Holaa!"
}
