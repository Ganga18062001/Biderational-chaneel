package main

import (
	"fmt"
	"sync"
)

//Bideractional channel used in sender and receiver
var ch chan int = make(chan int)

func Sender(gm *sync.WaitGroup) {
	defer gm.Done()
	ch <- 10
	fmt.Println("sending data completed")

}
func Receiver(gm *sync.WaitGroup){
	defer gm.Done()
	value := <-ch
	fmt.Println(value)
}

func main() {
	var gm sync.WaitGroup
	gm.Add(2)


	go Sender(&gm)
	go Receiver(&gm)
	gm.Wait()
	fmt.Println("ececution is completed")
	

}