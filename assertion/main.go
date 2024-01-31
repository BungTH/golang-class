package main

import "fmt"

func String(s interface{}) string {
	ss, ok := s.(string)
	if !ok {
		return ""
	}

	return ss
}

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	if !ok {
		fmt.Println("not ok")
	}
	fmt.Println(f, ok)

	// ff := i.(float64) // panic
	// fmt.Println(ff)

	fmt.Println(String("hello"))
	fmt.Println(String(123))
	fmt.Println(String(true))
}