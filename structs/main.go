package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Address struct {
	No string
}

type Employee struct {
	Name string
	Address
}

func main() {
	fmt.Println(Vertex{
		X: 1,
		Y: 2,
	})

	v := Vertex{}

	v.X = 3
	v.Y = 4

	fmt.Println(v)

	e := Employee{}

	e.Name = "John Wick"
	e.Address.No = "123"

	fmt.Println(e)
}