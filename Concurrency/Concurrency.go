package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("started")
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Sleep in 5")
	}()

	go func() {
		fmt.Println("Line")
	}()
}
