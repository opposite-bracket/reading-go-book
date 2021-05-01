package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Println("Program took", elapsed)
	}()

	var s, sep string
	for i, arg := range os.Args[1:] {
		s += fmt.Sprintf("%s%d %v", sep, i, arg)
		sep = "\n"
	}
	fmt.Println(s)
}
