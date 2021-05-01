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

	fmt.Println(os.Args[0])
}
