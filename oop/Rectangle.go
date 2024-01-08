package oop

type Rectangle struct {
	Common
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	// t := r.Common.GetName()
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}
func (c *Rectangle) GetName() string {
	return c.Name
}
