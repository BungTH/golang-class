package main

import "fmt"

var a int

const pi = 3.14

func main() {
	var b, c, d = true, false, "No!"

	e := 0

	fmt.Println(a, b, c, d, e)

	var f int = 13
	var g float64 = float64(f)
	var h string

	fmt.Printf("%v %v %q", f, g, h)	
}