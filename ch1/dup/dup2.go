package main

import (
	"bufio"
	"fmt"
	"os"
)

type blork struct {
	count    int
	filename string
}

func main() {
	counts := make(map[string]blork)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.count, n.filename, line)
		}
	}
}

func countLines(f *os.File, counts map[string]blork) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		var tmp = counts[input.Text()]
		tmp.count++
		tmp.filename = f.Name()
		counts[input.Text()] = tmp
	}
}
