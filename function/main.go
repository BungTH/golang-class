package main

import "fmt"

func adder(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y,x
}

func square(s int) (area, perimeter int) {
	area = s * s
	perimeter = 4 * s
	return
}

func main() {
	fmt.Println(adder(1, 2))

	fmt.Println(swap("world", "hello"))

	fmt.Println(square(2))
}