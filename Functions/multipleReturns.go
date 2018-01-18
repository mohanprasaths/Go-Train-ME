package main

import "fmt"

type Greetings struct {
	name          string
	greeting      string
	informalGreet string
}

func Printme(myGreetings Greetings) {
	formal, informal := GreetingsConvertor(myGreetings)
	fmt.Println(formal)
	fmt.Println(informal)
}

func GreetingsConvertor(myGreetings Greetings) (formalWish, informalWish string) {
	formalWish = myGreetings.greeting + "  " + myGreetings.name
	informalWish = myGreetings.informalGreet + "  " + myGreetings.name
	return
}

func main() {
	var myGreetings = Greetings{name: "mohan", greeting: "hello", informalGreet: "heeey"}
	Printme(myGreetings)
}
