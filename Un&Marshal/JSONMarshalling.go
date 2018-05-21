package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	ID      int      `json:"id"`
	Message string   `json:"mymessage"`
	Creator string   `json:"creator"`
	Tags    []string `json:"tags"`
}

func main() {
	sampleMessage := message{1, "hellos", "mohan", []string{"blog", "karate"}}
	marshalledStruct, err := json.Marshal(sampleMessage)
	if err != nil {
		fmt.Println("error")
	}
	var theNewMessage message
	fmt.Println(marshalledStruct)

	err = json.Unmarshal(marshalledStruct, &theNewMessage)
	if err != nil {
		fmt.Print("Erro")
	}
	fmt.Println(theNewMessage.Message)
	fmt.Println(sampleMessage)

	jsonString := `{"id":1,"mymessage":"hellosss","creator":"monger","tags":["blogge"]}`
	var theJSONNewMessage message
	err = json.Unmarshal([]byte(jsonString), &theJSONNewMessage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(theJSONNewMessage.Message)
}
