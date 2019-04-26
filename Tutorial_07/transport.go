package main

import(
    "fmt"
    "math/rand"
    "time"
)

// functions to generate time duration of various activities
func arrivalLapse() time.Duration {
    return time.Duration(rand.Intn(1000)) * time.Millisecond
}

func departureLapse() time.Duration {
    return time.Duration(rand.Intn(2000)) * time.Millisecond
}

func journeyDuration() time.Duration {
    return time.Duration(800) * time.Millisecond
}

func cableCarDuration() time.Duration {
    return time.Duration(500) * time.Millisecond
}

func operateLapse() time.Duration {
    return time.Duration(rand.Intn(3000)) * time.Millisecond
}

// the number of groups to send round the villages
const numGroups = 3

// producer generates groups and delivers them to the cablecar
func producer(t chan<- int) {
    for i:=0; i<numGroups; i++ {
        g := i+1
        fmt.Printf("group [%v] arrives on the mountain\n", g)
        t <- g
        time.Sleep(arrivalLapse())
    }
}

// consumer removes from the cablecar after their tour
func consumer(t <-chan int, done chan<- bool) {
    for i:=0; i<numGroups; i++ {
        g := <-t
        fmt.Printf("group [%v] departs from the mountain\n", g)
        time.Sleep(departureLapse())
    }
    done <- true
}

// train object, storing source and destination channels
type Train struct {
    id int
    source chan int
    destination chan int
}

// train method, handling transfer of groups from source to destination
func (t *Train) journey() {
    for {
        g := <-t.source
        fmt.Printf("train [%v] picks up group [%v]\n", t.id, g)
        time.Sleep(journeyDuration())
        t.destination <- g
        fmt.Printf("train [%v] drops off group [%v]\n", t.id, g)
        time.Sleep(journeyDuration())
    }
}

// cable car object, storing channels connecting 
type CableCar struct {
    atTop bool
    occupied bool
    lastTrain chan int
    firstTrain chan int
    arrival chan int
    departure chan int
}

// cable car method, handling transfer of groups up and down the mountain
func (cc *CableCar) travel() {
    for {
        if cc.atTop {
            g := <-cc.lastTrain
            fmt.Printf("group [%v] arrives at cable car (top)\n", g)
            cc.occupied = true
            fmt.Printf("cable car descending with group [%v]\n", g)
            time.Sleep(cableCarDuration())
            cc.atTop = false
            cc.departure <- g
            fmt.Printf("group [%v] departs from cable car (bottom)\n", g)
            cc.occupied = false
        } else {
            g := <-cc.arrival
            fmt.Printf("group [%v] arrives at cable car (bottom)\n", g)
            cc.occupied = true
            fmt.Printf("cable car ascending with group [%v]\n", g)
            time.Sleep(cableCarDuration())
            cc.atTop = true
            cc.firstTrain <- g
            fmt.Printf("group [%v] departs from cable car (top)\n", g)
            cc.occupied = false
        }
    }
}

// cable car method, handling movement of empty cable car
func (cc *CableCar) operate() {
    for {
        if !cc.occupied {
            if cc.atTop {
                fmt.Println("empty cable car descending")
                time.Sleep(cableCarDuration())
            } else {
                fmt.Println("empty cable car ascending")
                time.Sleep(cableCarDuration())
            }
            cc.atTop = !cc.atTop
            time.Sleep(operateLapse())
        }
    }
}

// number of trains to simulate (ie, 1 more than number of villages)
const numTrains = 3

func main() {
    done := make(chan bool) // channel to signal when simulatoin is done
    trains := make([]*Train, numTrains) // array of train channels

    // create first train, with both channels
    trains[0] = &Train{id: 0, 
		       source: make(chan int, 1), 
		       destination: make(chan int, 1)}

    // create rest of trains, daisy chained together
    for i:=1; i<numTrains; i++ {
        trains[i] = &Train{id: i, 
			   source: trains[i-1].destination, 
			   destination: make(chan int, 1)}
    }

    // create cable car, linking to first and last trains,
    // and creating two new channels to connect to producer/consumer
    cablecar := &CableCar{occupied: false, atTop: false, 
        firstTrain: trains[0].source, 
	lastTrain:trains[numTrains-1].destination,
        arrival: make(chan int), 
	departure: make(chan int)}

    // set producer & consumer goroutines going
    go producer(cablecar.arrival)
    go consumer(cablecar.departure, done)

    // set cablecar goroutines going
    go cablecar.travel()
    go cablecar.operate()

    // set all train goroutines going
    for i:=0; i<numTrains; i++ {
        go trains[i].journey()
    }

    // catch done signal
    <-done
}


