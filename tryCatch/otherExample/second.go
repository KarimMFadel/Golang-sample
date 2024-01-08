package main

import (
	"fmt"
)

func a() {
	defer func() {
		fmt.Println("defer a")
		if e := recover(); e != nil {
			fmt.Println("recovered", e)
		}
		fmt.Println("end defer of a")
	}()
	fmt.Println("a")
	b()
	fmt.Println("end a")
}

func b() {
	defer func() {
		fmt.Println("defer b")
		if e := recover(); e != nil {
			fmt.Println("recovered", e)
		}
		fmt.Println("end defer of b")
	}()
	fmt.Println("b")
	c()
	fmt.Println("end b")
}

func c() {
	defer func() {
		fmt.Println("defer c")
	}()
	fmt.Println("c")
	panic("c panic")
	fmt.Println("end c")
}

func main() {
	fmt.Println("start")
	a()
	fmt.Println("end")
}
