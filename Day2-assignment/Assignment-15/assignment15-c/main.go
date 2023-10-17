/*Now Modify the previous Go program to use three sender goroutines and two receiver
goroutines.
*/

package main

import (
	"fmt"
	"time"
)

func sender(ch chan int, senderID int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send numbers 1 to 5 to channel
		fmt.Printf("Sender %d sent: %d\n", senderID, i)
	}
}

func receiver(ch chan int, receiverID int) {
	for num := range ch { //sending value to receiver
		fmt.Printf("Receiver %d received: %d\n", receiverID, num) //printing the received numbers from channels
	}
}

func main() {

	ch := make(chan int) // Create unbuffered channel

	numSenders := 3
	numReceivers := 2

	for i := 1; i <= numSenders; i++ { // Start sender goroutines
		go sender(ch, i) // calling sender goroutine
	}

	for i := 1; i <= numReceivers; i++ { // Start receiver goroutines
		go receiver(ch, i) // calling receiver goroutine
	}

	time.Sleep(2 * time.Second) //calling time.sleep to pause the execution

	close(ch) // Close channel

	fmt.Println("End of the program")
}
