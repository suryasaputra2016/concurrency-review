package basicgoroutine

import (
	"fmt"
	"sync"
)

// BasicGoroutine call printNumbers and printLetters concurrently
// wait for them to finish before exiting
func BasicGoroutine() {
	var wg sync.WaitGroup

	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)
	wg.Wait()
}

// printNumbers print consecutive integers from 1 to 10
func printNumbers(wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

// printLetters print consecutive letters from a to j
func printLetters(wg *sync.WaitGroup) {
	for _, j := range "abcdefghij" {
		fmt.Printf("%c\n", j)

	}
	wg.Done()
}
