package main

import "fmt"

// geometry embeds both area and perimeter interfaces
type geometry interface {
	area
	perimeter
}

type area interface {
	area() float64
}

type perimeter interface {
	perimeter() float64 // Fixed typo from 'permiter'
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

// Added missing perimeter method for square
func (s square) perimeter() float64 {
	return 4 * s.side
}

type rectangle struct { // Fixed typo from 'rectange'
	width, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

// Added missing perimeter method for rectangle
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func calculateArea(s area) float64 {
	return s.area()
}

func cal(ge geometry) {
	fmt.Println("area:", ge.area())
	fmt.Println("perimeter:", ge.perimeter()) // Fixed typo
}

func main() {
	rect := rectangle{
		height: 12,
		width:  12,
	}
	sq := square{side: 12}

	fmt.Println("Shapes:", rect, sq)
	fmt.Println("Individual Areas:", calculateArea(rect))
	fmt.Println("Individual Areas:", calculateArea(sq))

	fmt.Println("\n--- Calling cal() for rect ---")
	cal(rect) // This works now!

	fmt.Println("\n--- Calling cal() for sq ---")
	cal(sq) // This works now too!
}
