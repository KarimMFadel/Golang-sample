package oop

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Common struct {
	Name string
}

func (c *Common) GetName() string {
	return c.Name
}
