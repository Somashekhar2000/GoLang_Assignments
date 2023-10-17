/* Simple concurrent program using go routine.
1. Create a Go program with two functions: sender and receiver.
2. The sender function should use a goroutine to send a sequence of numbers (e.g., 1
to 5) to a channel.
3. The receiver function should use a goroutine to receive numbers from the channel
and print each received number.
4. Use an unbuffered channel to ensure that the sender and receiver synchronize their
operations.
5. In the main function, start both the sender and receiver goroutines.
6. Use the time.Sleep function to allow the goroutines to execute. You can sleep for a
few seconds to ensure the sender and receiver have time to complete.(Using
time.Sleep in your code for synchronization is a bad practice but to keep this
example simple you can use time.Sleep here for now.)
7. Print a message in the main function to indicate the end of the program.
*/

package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send numbers 1 to 5 to channel
	}
	close(ch) // Close channel when done
}

func receiver(ch chan int) {
	for {
		num, ok := <-ch // Receive number from channel
		if !ok {
			break // Break loop when channel closed
		}
		fmt.Println("Received:", num) //printing the received numbers from channels
	}
}

func main() {
	ch := make(chan int) // Create unbuffered channel

	go sender(ch)   // creating sender goroutine
	go receiver(ch) //creating receiver goroutine

	time.Sleep(3 * time.Second) // Sleep for few seconds to allow sender and receiver goroutine to complete

	fmt.Println("End of the program")
}
