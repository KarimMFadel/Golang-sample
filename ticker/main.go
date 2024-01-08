package main

import (
	"fmt"
	"log"
	"time"
)

var (
	clear = make(chan bool)
)

func main() {
	testTimeTicker2()
	testTimeTicker()
}

func testTimeTicker2() {
	MyTicker1 := time.NewTicker(500 * time.Millisecond)
	MyTicker2 := time.NewTicker(1 * time.Second)

	go func() {
		for {
			<-MyTicker1.C
			log.Println("Tick Received for 500 millisecond")
		}
	}()

	go func() {
		for {
			<-MyTicker2.C
			log.Println("Tick Received 1 second")
		}
	}()

	time.Sleep(6 * time.Second)
	log.Println("Main finished")
	//select {}   // Because we use ‘select’ statement, the tickers will not be stopped and the program run infinitely
}

func testTimeTicker() {
	// A counter for the number of times we print
	printed := 0

	// We call set interval to print Hello World forever
	// every 1 second
	setInterval(func(num int) {
		time.Sleep(time.Duration(5-num) * time.Second)
		fmt.Printf("Hello World num:%d \n", num)
		printed++
	}, 1000, true)

	// If we wanted to we had a long running task (i.e. network call)
	// we could pass in true as the last argument to run the function
	// as a goroutine

	// Some artificial work here to wait till we've printed
	// 5 times
	for {
		if printed == 5 {
			// Stop the ticket, ending the interval go routine
			clear <- true
			return
		}
	}
}

func setInterval(someFunc func(num int), milliseconds int, async bool) chan bool {

	// How often to fire the passed in function
	// in milliseconds
	interval := time.Duration(milliseconds) * time.Millisecond

	// Setup the ticket and the channel to signal
	// the ending of the interval
	ticker := time.NewTicker(interval)
	//clear := make(chan bool)

	//defer ticker.Stop()

	// Put the selection in a go routine
	// so that the for loop is none blocking
	i := 1
	go func() {
		for {

			select {
			case <-ticker.C:
				if async {
					// This won't block
					go someFunc(i)
				} else {
					// This will block
					someFunc(i)
				}
			case <-clear:
				ticker.Stop()
				return
			}
			i++
		}
	}()

	// We return the channel so we can pass in
	// a value to it to clear the interval
	return clear
}
