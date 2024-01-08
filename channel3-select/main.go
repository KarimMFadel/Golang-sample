package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// create channels
	numberChan := make(chan int)
	// numberErrChan := make(chan error, 1)

	messageChan := make(chan string)
	// messageErrChan := make(chan error, 1)
	errChan := make(chan error, 1)

	var number int
	var text string
	var wg sync.WaitGroup

	wg.Add(2)
	go channelNumber(numberChan, errChan, &wg)
	go channelMessage(messageChan, errChan, &wg)

	fmt.Println("before select statements")

	var errN, errM error

	// text = <-messageChan
	// errM = <-errChan
	// if nil != errM {
	// 	fmt.Println("after receive an errMessage :", errM)
	// } else {
	// 	fmt.Printf("text: %s \n", text)
	// }

	// errN = <-errChan
	// number = <-numberChan
	// if nil != errN {
	// 	fmt.Println("after receive an errMessage2 :", errN)
	// } else {
	// 	fmt.Printf("number: %d \n", number)
	// }

	select {
	case errMessage := <-errChan:
		fmt.Println("after receive an errMessage2 :", errMessage)
		errM = errMessage

	case message := <-messageChan:
		fmt.Println("Channel Data message:", message)
		text = message
		fmt.Printf("text: %s \n", message)
	}

	select {
	case errNumber := <-errChan:
		fmt.Println("after receive an errMessage  :", errNumber)
		errN = errNumber

	case value := <-numberChan:
		fmt.Println("Channel Data number:", value)
		number = value
		fmt.Printf("number: %d \n", number)
	}

	wg.Wait()

	fmt.Printf("After select statements ======>>  text: %s, number: %d. \n", text, number)
	fmt.Printf("Errors ======>>  errMessage: %s, errNumber: %s. \n", errM, errN)

	fmt.Println("Before: final operation to check panic")
	time.Sleep(2 * time.Second)
	fmt.Println("After: final operation to check panic")
}

func channelNumber(number chan int, err chan error, wg *sync.WaitGroup) {
	defer close(err)
	defer close(number)
	defer wg.Done()

	fmt.Println("channelNumber method")
	time.Sleep(2 * time.Second)
	fmt.Println("channelNumber method second ==================")
	err <- fmt.Errorf("22f2 isn't number")
	// number <- 15
}

func channelMessage(message chan string, err chan error, wg *sync.WaitGroup) {
	defer close(err)
	defer close(message)
	defer wg.Done()

	fmt.Println("channelMessage method")
	time.Sleep(2 * time.Second)
	fmt.Println("channelNumber method second ==================")
	err <- fmt.Errorf("$$FF& isn/'t string")
	//message <- "Learning Go Select"
}
