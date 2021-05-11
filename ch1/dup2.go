package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			log.Print(f)
			countLines(f, count)
		}
	}
	// Ignoring input errors
	for s, i := range count {
		log.Printf("%d\t%s\n", i, s)
	}
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if "wq" == input.Text() {
			break
		}
		count[input.Text()]++
	}
}
