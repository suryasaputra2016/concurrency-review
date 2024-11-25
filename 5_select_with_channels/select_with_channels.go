package selectwithchannels

import "fmt"

// SelectWIthChannels print numbers 1 t 20 and also information wether they are odd of even
func SelectWithChannels() {
	odd := make(chan int)
	even := make(chan int)
	done := make(chan int)

	go sendNumbers(odd, even, done)

	for {
		select {
		case i := <-odd:
			fmt.Printf("received and odd number: %d \n", i)
		case i := <-even:
			fmt.Printf("received and even number: %d \n", i)
		case <-done:
			return
		}
	}
}

// sendNumbers send number from 1 to 20
// it sends odd number and even numbers to their respective channel
func sendNumbers(odd, even, done chan int) {
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			odd <- i
		} else {
			even <- i
		}
	}
	done <- 1
}
