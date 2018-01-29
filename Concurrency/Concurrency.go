package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func additionExpert(a int, b int) (sum int) {
	sum = a + b
	time.Sleep(7 * time.Second)
	fmt.Println(sum)
	mySync <- true
	return
}

var mySync = make(chan bool, 1)

func main() {
	runtime.GOMAXPROCS(1)
	var WaitGrp sync.WaitGroup
	WaitGrp.Add(2)
	fmt.Println("started")
	go func() {
		defer WaitGrp.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Sleep and run in 2 - thread1")
	}()
	go func() {
		defer WaitGrp.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Sleep and run in 2 - thread 2")
	}()
	go func() {
		defer WaitGrp.Done()
		fmt.Println("Priority run")
	}()
	WaitGrp.Wait()

	go additionExpert(3, 7)
	<-mySync
}
