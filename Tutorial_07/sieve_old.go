// The Sieve of Erasthones

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const KSIEVEOLD = 20  // number of primes to generate

// generate a sequence of numbers beginning from 2
func producerSieveOld(ch chan<- int, chLength chan<- int) {
	// TODO
	length := 0
	for i := 2; i <= KSIEVEOLD; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		ch <- i
		length += 1
	}
	chLength <- length
}

// read values from channel in to channel out
// removing those divisible by 'prime'
func filterOld(in <-chan int, out chan<- int, prime int, inLength int, chLength chan<- int) {
	// TODO
	length := 0
	for i := 0; i < inLength - 1; i++ {
		inVar := <-in
		if 0 != (inVar % prime) {
			out <- inVar
			length += 1
		}
	}
	chLength <- length
}

// Sieve of Erasthones: daisy-chain filter processes
func main() {
	// TODO
	ch := make(chan int, KSIEVE)
	chLength := make(chan int, 1)
	go producerSieveOld(ch, chLength)
	in := ch
	length :=<-chLength
	for length != 0 {
		prime :=<-in
		fmt.Println(prime)
		out := make(chan int, length)
		filterOld(in, out, prime, length, chLength)
		in = out
		length =<-chLength
	}
}
