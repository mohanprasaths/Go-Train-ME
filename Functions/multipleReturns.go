package main

import "fmt"

type Greetings struct {
	name          string
	greeting      string
	informalGreet string
}

type Printer func(string)

var greetingMap map[string]string

func Printme(myGreetings Greetings, do Printer) {
	formal, informal := GreetingsConvertor(myGreetings)
	fmt.Println(formal)
	fmt.Println(informal)
	formal, informal, variaDic := variaDicGreetingsConvertor(myGreetings, "working", "  fine")
	do(formal)
	do(informal)
	do(variaDic)
}

func PrintmeSlice(myGreetings []Greetings, do Printer) {
	for _, s := range myGreetings {
		formal, informal := GreetingsConvertor(s)
		fmt.Println(formal)
		fmt.Println(informal)
		formal, informal, variaDic := variaDicGreetingsConvertor(s, "working", "  fine")
		do(formal)
		do(informal)
		do(variaDic)
	}
}

func GreetingsConvertor(myGreetings Greetings) (formalWish, informalWish string) {
	formalWish = myGreetings.greeting + "  " + myGreetings.name
	informalWish = myGreetings.informalGreet + "  " + myGreetings.name
	return
}

func variaDicGreetingsConvertor(myGreetings Greetings, variaDic ...string) (formalWish, informalWish, variaDicEg string) {
	formalWish = myGreetings.greeting + "  " + getPrefixUsingMap(myGreetings.name) + myGreetings.name
	informalWish = myGreetings.informalGreet + "  " + myGreetings.name
	variaDicEg = variaDic[0] + variaDic[1]
	return
}

func getPrefixUsingMap(name string) (prefix string) {
	prefix = greetingMap[name]
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

func initGreetingMap() {
	greetingMap = make(map[string]string)
	greetingMap["mohan"] = "Mr"
	greetingMap["prasath"] = "Dr"
	greetingMap["Superman"] = "hero"
	greetingMap["Spiderman"] = "hero"
	greetingMap["Ironman"] = "Boss"

	delete(greetingMap, "Ironman")
}

func initSlice() {
	slice := []int{1, 2, 3, 4, 5}
	slice = appendSliceWithSlice(slice, 6)
	slice = appendSlice(slice, 7)

	fmt.Print(slice)
}

func appendSlice(slice []int, number int) (sliceNew []int) {
	sliceNew = append(slice, number)
	return
}
func appendSliceWithSlice(slice []int, number int) (sliceNew []int) {
	sliceTemp := []int{number}
	sliceNew = append(slice, sliceTemp...)
	return
}
func main() {
	initSlice()
	initGreetingMap()
	var myGreetings = Greetings{name: "mohan", greeting: "hello", informalGreet: "heeey"}
	Printme(myGreetings, myPrint)
	Printme(myGreetings, myPrintLine)
	Printme(myGreetings, myCustomPrint("!!!!"))
	var myGreetings2 = Greetings{name: "prasath", greeting: "hello", informalGreet: "heeey"}
	Printme(myGreetings2, myPrint)
	Printme(myGreetings2, myPrintLine)
	Printme(myGreetings2, myCustomPrint("!!!!"))

	slice := []Greetings{
		{name: "Spiderman", greeting: "Superhero", informalGreet: "Bugg"},
		{name: "Superman", greeting: "Superhero", informalGreet: "Alien"},
		{name: "Ironman", greeting: "Superhero", informalGreet: "Tin"},
	}

	PrintmeSlice(slice, myPrint)
	PrintmeSlice(slice, myPrintLine)

}
