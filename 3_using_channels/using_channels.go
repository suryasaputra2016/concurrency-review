package usingchannels

import "fmt"

func UsingChannels() {
	stream := make(chan int)
	done := make(chan int)

	go producer(stream)
	go consumer(stream, done)

	<-done
}

// producer generates consecutive integers from 1 to 10
// send them to stream
func producer(stream chan int) {
	for i := 1; i <= 10; i++ {
		stream <- i
	}
}

// consumer read integers from stream and print it out
func consumer(stream, done chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(<-stream)
	}
	done <- 1
}
