package main

import (
	"fmt"
	//"runtime"
)

func isOdd(n int) bool {
	// if n % 2 == 0 {
	// 	return false
	// }
	// return true
	return n % 2 != 0
}

func main() {
	// sum := 1

	// for i := 0; i <= 10; i++ {
	// 	sum += i
	// }

	// fmt.Println(sum)

	// for {
	// 	fmt.Println(sum)
	// }

	// fmt.Println(isOdd(2))
	// fmt.Println(isOdd(3))

	// switch os := runtime.GOOS; os {
	// 	case "darwin":
	// 		fmt.Println("MacOS")
	// 	case "windows":
	// 		fmt.Println("Windows")
	// 	default:
	// 		fmt.Println("Unknown")
	// }

	defer fmt.Println("World")
	fmt.Println("Hello")
}