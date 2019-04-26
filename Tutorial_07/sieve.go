// The Sieve of Erasthones

package main

import (
	"fmt"
)

const KSIEVE = 20  // number of primes to generate

// generate a sequence of numbers beginning from 2
func producerSieve(ch chan<- int) {
	// TODO
	for i := 2; ; i++ {
		ch <- i
	}
}

// read values from channel in to channel out
// removing those divisible by 'prime'
func filter(in <-chan int, out chan<- int, prime int) {
	// TODO
	for {
		inVar := <-in
		if 0 != (inVar % prime) {
			out <- inVar
		}
	}
}

// Sieve of Erasthones: daisy-chain filter processes
func main() {
	// TODO
	ch := make(chan int, KSIEVE)
	go producerSieve(ch)
	in := ch
	for i := 0; i < KSIEVE; i++ {
		prime :=<-in
		out := make(chan int)
		go filter(in, out, prime)
		in = out
		fmt.Printf("%d ", prime)
	}
}
