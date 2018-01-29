package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var WaitGrp sync.WaitGroup
	WaitGrp.Add(2)
	fmt.Println("started")
	go func() {
		defer WaitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Sleep and run in 5")
	}()

	go func() {
		defer WaitGrp.Done()
		fmt.Println("Priority run")
	}()
	WaitGrp.Wait()
}
