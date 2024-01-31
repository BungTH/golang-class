package main

import (
	"fmt"

	"github.com/bungkapth/exercise2/grading"
)

func main() {
	fmt.Println(grading.Grade(101))
	fmt.Println(grading.Grade(80))
	fmt.Println(grading.Grade(70))
	fmt.Println(grading.Grade(60))
	fmt.Println(grading.Grade(50))
	fmt.Println(grading.Grade(49))
}