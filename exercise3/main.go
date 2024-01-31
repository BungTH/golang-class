package main

import "fmt"

func double(pointer *int) {
	println(pointer)
	println(*pointer)
	*pointer = *pointer * 2

}

func main() {
	n := 10

	fmt.Println(n)

	double(&n)

	fmt.Println(n)
}