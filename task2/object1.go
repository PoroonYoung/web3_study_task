package main

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct{}

func (r *Rectangle) Area() {
	fmt.Println("Rectangle类型对象调用Area()")
}
func (r *Rectangle) Perimeter() {
	fmt.Println("Rectangle类型对象调用Perimeter()")
}

type Circle struct{}

func (c *Circle) Area() {
	fmt.Println("Circle类型对象调用Area()")
}
func (c *Circle) Perimeter() {
	fmt.Println("Circle类型对象调用Perimeter()")
}

func main() {
	rectangle := Rectangle{}
	rectangle.Area()
	rectangle.Perimeter()
	circle := Circle{}
	circle.Area()
	circle.Perimeter()
}
