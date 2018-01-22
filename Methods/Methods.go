package main

import (
	"fmt"
)

type Greetings struct {
	name          string
	greeting      string
	informalGreet string
}

type GreetingsCollection []Greetings

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
		// fmt.Println(formal)
		// fmt.Println(informal)
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

func (greetingsColl GreetingsCollection) GreetUS(do Printer) {
	for _, s := range greetingsColl {
		formal, informal, variaDic := variaDicGreetingsConvertor(s, "Greetings", "  Collection using Method")
		do(formal)
		do(informal)
		do(variaDic)
	}
}

func myPrinter(printString string) {
	fmt.Println(printString + " ->>>>>> Go routine")
}

func (greeting *Greetings) Rename(newName string) {
	greeting.name = newName
}

type Renamable interface {
	Rename(newString string)
}

func RenameToIronman(r Renamable) {
	r.Rename("Ironman")
}

func (greetingsColl GreetingsCollection) GreetingsCollectionFiller(c chan Greetings) {
	for _, s := range greetingsColl {
		c <- s
	}
	close(c)
}

func main() {
	slice := GreetingsCollection{
		{name: "Spiderman", greeting: "Superhero", informalGreet: "Bugg"},
		{name: "Ironman", greeting: "Superhero", informalGreet: "Tinman"},
	}

	// PrintmeSlice(slice, myPrintLine)
	done := make(chan bool, 2)
	go func() {
		slice.GreetUS(myPrinter)
		done <- true
	}()
	slice.GreetUS(myPrintLine)
	<-done

	c := make(chan Greetings)
	go slice.GreetingsCollectionFiller(c)
	for s := range c {
		fmt.Println(s.name)
	}
}
