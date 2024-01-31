package main

import "fmt"

func main() {
	i, j := 20, 50

	p := &i         // point to i

	fmt.Println(p)  // print address of i
	fmt.Println(*p) // print value of i

	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j          // point to j
	fmt.Println(p)  // print address of j
	fmt.Println(*p) // print value of j

	*p = *p / 5     // divide j through the pointer
	fmt.Println(j)  // see the new value of j
}