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
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
