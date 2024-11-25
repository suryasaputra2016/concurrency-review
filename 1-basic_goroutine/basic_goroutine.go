package basicgoroutine

import (
	"fmt"
)

// BasicGoroutine call printNumbers and printLetters concurrently
func BasicGoroutine() {
	go printNumbers()
	go printLetters()
}

// printNumbers print consecutive integers from 1 to 10
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// printLetters print consecutive letters from a to j
func printLetters() {
	for _, j := range "abcdefghij" {
		fmt.Printf("%c\n", j)

	}
}
