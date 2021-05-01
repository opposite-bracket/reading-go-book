package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Println("Program took", elapsed)
	}()

	fmt.Println(strings.Join(os.Args[1:], " "))
}
