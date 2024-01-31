package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MyFloat is a type that we will use to implement the Abser interface
type MyFloat float64

// Abs is a method on the type MyFloat which calculates the absolute value of a MyFloat
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f) // if f is less than 0, return its negation (making it positive)
    }
    return float64(f) // if f is 0 or positive, return it as is
}

func main() {
    var a Abser // declare a variable 'a' of interface type Abser
    f  := MyFloat(-math.Sqrt(2)) // initialize 'f' as a MyFloat with value -sqrt(2)
    v  := Vertex{3, 4} // initialize 'v' as a Vertex with x=3 and y=4
    a = f // assign 'f' to 'a'. This is possible because MyFloat implements Abser
    a = &v // assign the address of 'v' to 'a'. This is possible because *Vertex (pointer to Vertex) implements Abser
    // a = v // this line is commented out because Vertex (not pointer to Vertex) does not implement Abser
    fmt.Println(a.Abs()) // call the Abs method on 'a', which will call the Abs method of the underlying type (Vertex in this case)
}