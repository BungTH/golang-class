package main

import (
	"fmt"

	"github.com/bungkapth/exercise5/word"
)

func main() {
	var count map[string]int = word.WordCount("word word hello hi word")

	fmt.Println(count)
}