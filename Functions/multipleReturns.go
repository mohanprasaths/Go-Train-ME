package main

import "fmt"

type Greetings struct {
	name          string
	greeting      string
	informalGreet string
}

type Printer func(string)

func Printme(myGreetings Greetings, do Printer) {
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
	formalWish = myGreetings.greeting + "  " + getPrefix(myGreetings.name) + myGreetings.name
	informalWish = myGreetings.informalGreet + "  " + myGreetings.name
	variaDicEg = variaDic[0] + variaDic[1]
	return
}

func getPrefix(name string) (prefix string) {
	switch name {
	case "mohan":
		prefix = "Mr "
	case "prasath":
		prefix = "Dr "
	default:
		prefix = "Dude"

	}
	return prefix
}

func myPrint(mystring string) {
	fmt.Print(mystring)
}
func myPrintLine(mystring string) {
	fmt.Println(mystring)
}

func myCustomPrint(custom string) Printer {
	return func(myString string) {
		fmt.Print(myString + custom)
	}
}

func main() {
	var myGreetings = Greetings{name: "mohan", greeting: "hello", informalGreet: "heeey"}
	Printme(myGreetings, myPrint)
	Printme(myGreetings, myPrintLine)
	Printme(myGreetings, myCustomPrint("!!!!"))
	var myGreetings2 = Greetings{name: "prasath", greeting: "hello", informalGreet: "heeey"}
	Printme(myGreetings2, myPrint)
	Printme(myGreetings2, myPrintLine)
	Printme(myGreetings2, myCustomPrint("!!!!"))
}
