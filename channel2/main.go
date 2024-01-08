package main

import (
	"fmt"
	"strings"
	"time"
)

func capsAndLen(words []string, cs chan string, ci chan int) {
	defer close(cs)
	defer close(ci)
	for _, word := range words {
		cs <- strings.ToUpper(word)
		ci <- len(word)
	}
}

func main() {
	// words := []string{"lorem", "ipsum", "dolor", "sit", "amet"}
	// cs := make(chan string)
	// ci := make(chan int)
	// go capsAndLen(words, cs, ci)
	// for allCaps := range cs {
	// 	length := <-ci
	// 	fmt.Println(allCaps, ",", length)
	// }

	// var wg sync.WaitGroup
	result := make(chan int, 1)
	err := make(chan error, 1)
	result2 := make(chan int, 1)
	err2 := make(chan error, 1)

	// wg.Add(2)
	// go testTwoChan(0, result, err, &wg)
	// go testTwoChan(10, result2, err2, &wg)
	// wg.Wait()

	go testTwoChan(0, result, err)
	go testTwoChan(0, result2, err2)

	res := <-result
	errorx := <-err
	if nil != errorx {
		fmt.Println("errrrrrrrrrrrrrrrrrror")
	} else {
		fmt.Println("X value 1 = ", res)
	}

	errorx2 := <-err2
	res2 := <-result2
	if nil != errorx2 {
		fmt.Println("errrrrrrrrrrrrrrrrrror2")
	} else {
		fmt.Println("X value 2 = ", res2)
	}

	/**
	 * when use select with the 2 error channels.
	 * the second select (sometime) didn't work correctly,
	 * and i saw on the log fmt.Println("X value 1 = ", res)
	 *
	 * To avoid this issue, so i make the if condition, like the abbove code.
	 **/

	// select {
	// case errorx := <-err:
	//     fmt.Println("errrrrrrrrrrrrrrrrrror", errorx)
	// case res := <-result:
	//     fmt.Println("X value 1 = ", res)
	// }
	// select {
	// case errorx := <-err2:
	//     fmt.Println("errrrrrrrrrrrrrrrrrror2", errorx)
	// case res := <-result2:
	//     fmt.Println("X value 1 = ", res)
	// }
}

// func testTwoChan(x int, result chan int, err chan error, wg *sync.WaitGroup) {
func testTwoChan(x int, result chan int, err chan error) {
	defer close(result)
	defer close(err)

	time.Sleep(time.Second * 2)

	if x != 0 {
		result <- 1000
		return
	}

	err <- fmt.Errorf("x value equals zero")
}
