// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	t1 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Now().Sub(t1))
}
