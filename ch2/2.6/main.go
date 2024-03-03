package main

import (
	"fmt"

	"tempconv.go/tempconv"
)

func main() {

	var f tempconv.Fahrenheit = 40

	c := tempconv.FToC(f)
	fmt.Println(c.String())

	k := tempconv.FToK(f)
	fmt.Println(k.String())
	c = tempconv.KToC(k)
	fmt.Println(c.String())
}
