package main

import "fmt"

type Square struct {
	X, Y int
}

type Rectangle struct {
	X, Y int
}

// func Area(x, y int) int {
// 	return x * y
// }

func (s Square) Area() int {
	return s.X * s.Y
}

func (s Square) ToString() string {
	return fmt.Sprintf("X: %d, Y: %d", s.X, s.Y)
}

func (s Square) MapToRectangle() Rectangle {
	return Rectangle(s)
}

func main() {

	sq := Square{10, 10}

	// fmt.Println(Area(10, 10))

	fmt.Println(sq.Area())

	fmt.Println(sq.ToString())

	fmt.Println(sq.MapToRectangle())
}