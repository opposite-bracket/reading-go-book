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

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
