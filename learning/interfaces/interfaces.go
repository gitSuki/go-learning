package main

import "fmt"

// here we are defining a brand new type called bot.
type bot interface {
	// this says that any type in the program with a function called getGreeting that returns a string
	// is now an honorary member of type bot. Now that they are honorary members of type bot, they can
	// access all functions connected to the bot type interface
	// interfaces don't have any direct utility on their own, we can't create an actual variable of
	// type bot, it's just useful for sharing functionality among pre-existing structs
	getGreeting() string
}

type englishBot struct{}
type russianBot struct{}

func main() {
	eb := englishBot{}
	rb := russianBot{}

	printGreeting(eb)
	printGreeting(rb)
}

func printGreeting(b bot) {
	// this is a function connected to the bot interface
	fmt.Println(b.getGreeting())
}

// having interfaces are very useful in this case where we have two structs which might
// have some areas of differing functionality, like calculating the greeting between
// english and russian will require different logic. but will also require some shared
// logic, printing the output will be the same regardless of the language of the bot.
func (englishBot) getGreeting() string {
	return "Hello"
}

// these technically aren't methods, we say that getGreeting expects russianBot as it's reciever
func (russianBot) getGreeting() string {
	return "Zdrastvuyte"
}
