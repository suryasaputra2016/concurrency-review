package handlingerrors

import "fmt"

// SelectWIthChannels print numbers 1 t 20 and also information wether they are odd of even
func HandlingErrors() {
	odd := make(chan int)
	even := make(chan int)
	err := make(chan int)
	done := make(chan int)

	go sendNumbers(odd, even, err, done)

	for {
		select {
		case i := <-odd:
			fmt.Printf("Received and odd number: %d \n", i)
		case i := <-even:
			fmt.Printf("Received and even number: %d \n", i)
		case <-done:
			return
		}
	}
}

// sendNumbers send number from 1 to 20
// it sends odd number and even numbers to their respective channel
func sendNumbers(odd, even, err, done chan int) {
	for i := 1; i <= 20; i++ {
		if i > 20 {
			err <- i
		} else if i%2 == 1 {
			odd <- i
		} else {
			even <- i
		}
	}
	done <- 1
}
