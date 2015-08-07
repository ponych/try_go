package main

import (
	"fmt"
	"math"
)


type Rectangle struct{
	height float64
	width float64
}

type Circle struct {
	 radius float64
}

func (c Circle)area() float64 {
	return c.radius * c.radius * math.Pi / 2.0
}

func (r Rectangle)area() float64 {
	return r.height * r.width
}

func main() {
	r := Rectangle{ 120, 20}
	c := Circle{3.19}

	fmt.Printf("area of rectangle is %v\n", r.area())
	fmt.Printf("area of circle is %v\n", c.area())
}