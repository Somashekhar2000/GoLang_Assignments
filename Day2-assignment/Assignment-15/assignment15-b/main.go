/*Now modify the previous code to use the buffered channel & run it again.*/

package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send numbers 1 to 5 to buffered channel
	}
	close(ch) // Close the channel when done
}

func receiver(ch chan int) {
	for num := range ch { //sending value to receiver
		fmt.Println("Received:", num) //printing the received numbers from channels
	}
}

func main() {
	ch := make(chan int, 5) // Create buffered channel with capacity 5

	go sender(ch)   // calling sender goroutine
	go receiver(ch) // calling receiver goroutine

	time.Sleep(3 * time.Second) // Sleep for few seconds to allow sender and receiver goroutines to complete

	fmt.Println("End of the program")
}
