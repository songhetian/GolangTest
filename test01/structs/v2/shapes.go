package v2

import "math"

type Rectangle struct {
	width, height float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {

	return r.width * r.height
}

func (c *Circle) Area() float64 {

	return math.Pi * c.Radius * c.Radius
}

//计算周长
func Perimeter(rectangle Rectangle) (p float64) {

	p = (rectangle.height + rectangle.width) * 2

	return
}

//计算面积

func Area(rectangle Rectangle) (p float64) {

	p = rectangle.width * rectangle.height

	return
}
