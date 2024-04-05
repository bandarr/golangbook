package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := map[string]int{
		"gilligan":  21,
		"skipper":   40,
		"professor": 35,
		"ginger":    29,
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
