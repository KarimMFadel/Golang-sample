package main

// source ==> https://dzone.com/articles/try-and-catch-in-golang

import (
	"errors"
	"fmt"
)

type Block[T any] struct {
	Try     func() T
	Catch   func(GenericError)
	Finally func()
}

func Throw(up GenericError) {
	panic(up)
}

func (tcf Block[T]) Do() T {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r.(GenericError))
			}
		}()
	}
	return tcf.Try()
}

func main() {
	value := doSomethings()
	fmt.Println("We went on ... " + value)
}

func doSomethings() string {
	fmt.Println("We started")
	return Block[string]{
		Try: func() string {
			fmt.Println("I tried")
			doSubTask()
			return "eeeeeee"
			// fmt.Println("After Throw") // unreachable code
		},
		Catch: func(e GenericError) {
			fmt.Printf("Caught %v\n", e.Error)
			//return "eeeeeee"
			fmt.Printf("Caught %v\n", e.Error)
		},
		Finally: func() {
			fmt.Println("Finally...")
		},
	}.Do()
	// fmt.Println("We went on")
}

func doSubTask() {
	Throw(GenericError{
		Error:  errors.New("sdsdsddssdsd"),
		Status: 303,
	})
}

type GenericError struct {
	Error  error
	Status int
}
