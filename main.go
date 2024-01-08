package main

import (
	"fmt"
	"oop/oop"
)

func main() {
	tryOop()
}

func tryOop() {
	r := oop.Rectangle{Common: oop.Common{Name: "Rectangle"}, Width: 5, Height: 10}
	fmt.Println("Name:", r.GetName())
	fmt.Println("Rectangle Area:", r.Area())
	fmt.Println("Rectangle Perimeter:", r.Perimeter())

	c := oop.Circle{Common: oop.Common{Name: "Circle"}, Radius: 7.5}
	fmt.Println("Name:", c.GetName())
	fmt.Println("Circle Area:", c.Area())
	fmt.Println("Circle Perimeter:", c.Perimeter())
}
