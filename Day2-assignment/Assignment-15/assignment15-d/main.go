package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() //decrementing count
	ch <- 2
	fmt.Println("sended : 2")
}

func receiver(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()               //decrementing count
	x := <-ch                     //sending value to receiver
	fmt.Println("received : ", x) //printing the received numbers from channels
}

func main() {
	ch := make(chan int)    // Create unbuffered channel
	wg := &sync.WaitGroup{} //creating waitgroup

	wg.Add(2)           //add number og goroutine to waitgroup
	go sender(ch, wg)   // calling sender goroutine
	go receiver(ch, wg) // calling receiver goroutine

	wg.Wait() // Wait for all sender and receiver goroutines to finish

	close(ch) // Close the channel when all senders are done

	fmt.Println("Program has finished.")
}
