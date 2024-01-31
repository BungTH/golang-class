package main

import "fmt"

func main() {
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a)

	var s []string

	fmt.Println(s)

	if s == nil {
		fmt.Println("Is nil!")
	}

	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)

	fmt.Println(primes[1:4])

	fmt.Println(primes[:4])

	fmt.Println(primes[1:])
}