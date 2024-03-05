package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"tempconv.go/tempconv"
)

func main() {

	var args []string = os.Args
	reader := bufio.NewReader(os.Stdin)

	if len(os.Args) == 1 {
		fmt.Print("Enter value: ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		args = append(args, text)
	}

	for _, arg := range args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
