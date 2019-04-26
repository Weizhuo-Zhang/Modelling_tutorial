// Conway's problem 

package main

import (
    "fmt"
    "strconv"
    "math/rand"
    "time"
)

const N = 100  // length of input character string
const MAX = 9   // max number of repeats to allow
const K = 40    // line length for output string

// function to compress a sequence of repeated characters
// received from inC and send result to outC
// eg "aaaaa" --> "5a"
func compress(inC <-chan string, pipe chan<- string) {
    c := ""  // current character
    previous := "" // previous character
    n := 0  // counter for repeats of current character
    previous = <-inC  // get first character
    for {

        // TODO

    }
}

// function to format a compressed sequence of characters
// printing K to a line
func output(pipe <-chan string, outC chan<- string) {
    c := ""  // current character
    n := 0  // counter for characters on current line
    for {

        // TODO

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
