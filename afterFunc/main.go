package main

import (
	"fmt"
	"time"
)

func main() {

	current_time := time.Now().UTC()
	fmt.Println(current_time.Format(time.RFC1123))

	testTimeAfterFunc2()
}
func testTimeAfterFunc() {
	ch := make(chan int)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("5 seconds passed")
		ch <- 10
	})
	for {
		select {
		case i := <-ch:
			fmt.Println(i, " is coming, end.")
			return
		default:
			fmt.Println("wait")
			time.Sleep(1 * time.Second)
		}
	}
}

func testTimeAfterFunc2() {
	printed := false
	print := func() {
		fmt.Println("This will print after x milliseconds")
		printed = true
	}

	// Make the timeout print after 5 seconds
	setTimeout(print, 5000)

	fmt.Println("This will print straight away")

	// Wait until it's printed our function string
	// before we close the program
	for {
		if printed {
			return
		}
	}

}
func setTimeout(someFunc func(), milliseconds int) {

	timeout := time.Duration(milliseconds) * time.Millisecond

	// This spawns a goroutine and therefore does not block
	time.AfterFunc(timeout, someFunc)

}
