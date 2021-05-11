package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if "wq" == input.Text() {
			break
		}
		count[input.Text()]++
	}
	// Ignoring input errors
	for s, i := range count {
		log.Printf("%d\t%s\n", i, s)
	}
}
