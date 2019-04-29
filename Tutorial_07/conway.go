// Conway's problem 

package main

import (
    "fmt"
    "strconv"
    "math/rand"
    "time"
)

const N = 10  // length of input character string
const MAX = 2   // max number of repeats to allow
const K = 40    // line length for output string
const ZERO = 0
const ONE = 1

// function to compress a sequence of repeated characters
// received from inC and send result to outC
// eg "aaaaa" --> "5a"
func compress(inC <-chan string, pipe chan<- string) {
    c := ""  // current character
    previous := "" // previous character
    n := ZERO // counter for repeats of current character
    previous = <-inC  // get first character
    for i := ONE; i < N;i++{
        // TODO
        c = <-inC
        n++
        if c != previous {
            if ONE == n {
                pipe <- previous
            } else {
                pipe <- strconv.Itoa(n)
                pipe <- previous
            }
           	n = ZERO
        } else {
            if MAX == n {
                pipe <- strconv.Itoa(n)
                pipe <- previous
                n = ZERO
            }
        }
        previous = c
    }
    pipe <- strconv.Itoa(n+1)
    pipe <- previous
}

// function to format a compressed sequence of characters
// printing K to a line
func output(pipe <-chan string, outC chan<- string) {
    c := ""  // current character
    n := ZERO  // counter for characters on current line
    for {
        // TODO
        c = <- pipe
        if K == n {
            outC <- "\n"
            n = ZERO
        }
        n++
        outC <- c
    }
}

// function to generate a random stream of a and b characters
func producer(inC chan<- string) {
    var letters = string("ab")
    for i:=0; i<N; i++ {
        inC <- string(letters[rand.Intn(len(letters))])
    }
}

// takes output and prints to screen
// uses a timeout to break loop when no further characters received
func consumer(outC <-chan string, done chan<- bool) {
    // timeout; sends a signal after 1 second
    timeout := make(chan bool, 1)
    go func() {
        time.Sleep(1 * time.Second)
        timeout <- true
    }()

    loop:  // label to indicate where we want to break to
    for {
        // either print character, if available
        // or break out of loop if timeout received
        select {
        case c := <-outC:
            fmt.Printf("%v", string(c))
        case <- timeout:
            break loop 
        }
    }
    done <- true
}

func main() {
    // create our channels
    inC := make(chan string)
    outC := make(chan string)
    pipe := make(chan string)
    done := make(chan bool)

    // set our goroutines running
    go producer(inC)
    go compress(inC, pipe)
    go output(pipe, outC)
    go consumer(outC, done)

    <-done
    print("\n")  // terminate nicely
}
