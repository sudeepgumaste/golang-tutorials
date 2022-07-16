package main

import "fmt"

type bot interface {
	getGreeting() string
}

func printGreeting (b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	var eb englishBot
	var sb spanishBot
	
	printGreeting(eb)
	printGreeting(sb)
}

func (e englishBot) getGreeting() string {
	return "Hi there"
}
func (s spanishBot) getGreeting() string {
	return "Hola!"
}