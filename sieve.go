// A concurrent prime sieve.
// This code is used for benchmarks and it should be considered as internal.
// See also https://golang.org/doc/play/sieve.go for original code.

package primes

import (
	"fmt"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func Sieve(num int) bool {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < num; i++ {
		prime := <-ch

		if prime > num {
			// fmt.Println(num, "is not prime")
			return false
		}

		if prime == num {
			// fmt.Println(num, "is prime")
			return true
		}

		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

	fmt.Println("sieve should not print this")
	return false
}
