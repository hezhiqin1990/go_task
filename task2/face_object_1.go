package main

import "fmt"

type Shape interface {
	Area()

	Perimeter()
}

type Rectangle struct {
	Name string
}

type Circle struct {
	Name string
}

func (r Rectangle) Area() {
	fmt.Println("Rectangle Area")
}

func (r Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter")
}

func (r Circle) Area() {
	fmt.Println("Circle Area")
}

func (r Circle) Perimeter() {
	fmt.Println("Circle Perimeter")
}

func main() {
	Rectangle := Rectangle{
		Name: "矩形",
	}
	var s1 Shape = Rectangle
	s1.Area()
	s1.Perimeter()
	var s2 Shape = Circle{
		Name: "圆圈",
	}
	s2.Area()
	s2.Perimeter()
}
