package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type blork struct {
	count     int
	filenames []string
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
			fmt.Printf("%d\t%s\t%s\n", n.count, strings.Join(n.filenames, ", "), line)
		}
	}
}

func countLines(f *os.File, counts map[string]blork) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		tmp := counts[input.Text()]
		tmp.count++
		if !contains(tmp.filenames, f.Name()) {
			tmp.filenames = append(tmp.filenames, f.Name())
		}
		counts[input.Text()] = tmp
	}
}

func contains(filenames []string, filename string) bool {
	for _, blah := range filenames {
		if blah == filename {
			return true
		}
	}
	return false
}
