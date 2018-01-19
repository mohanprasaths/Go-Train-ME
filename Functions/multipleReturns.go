package main

import "fmt"

type Greetings struct {
	name          string
	greeting      string
	informalGreet string
}

func Printme(myGreetings Greetings, do func(string)) {
	formal, informal := GreetingsConvertor(myGreetings)
	fmt.Println(formal)
	fmt.Println(informal)
	formal, informal, variaDic := variaDicGreetingsConvertor(myGreetings, "working", "  fine")
	do(formal)
	do(informal)
	do(variaDic)
}

func GreetingsConvertor(myGreetings Greetings) (formalWish, informalWish string) {
	formalWish = myGreetings.greeting + "  " + myGreetings.name
	informalWish = myGreetings.informalGreet + "  " + myGreetings.name
	return
}

func variaDicGreetingsConvertor(myGreetings Greetings, variaDic ...string) (formalWish, informalWish, variaDicEg string) {
	formalWish = myGreetings.greeting + "  " + myGreetings.name
	informalWish = myGreetings.informalGreet + "  " + myGreetings.name
	variaDicEg = variaDic[0] + variaDic[1]
	return
}

func myPrint(mystring string) {
	fmt.Print(mystring)
}
func myPrintLine(mystring string) {
	fmt.Println(mystring)
}

func main() {
	var myGreetings = Greetings{name: "mohan", greeting: "hello", informalGreet: "heeey"}
	Printme(myGreetings, myPrint)
	Printme(myGreetings, myPrintLine)
}
