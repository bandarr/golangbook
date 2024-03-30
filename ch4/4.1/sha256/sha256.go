package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("q"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	blorg := 0

	for j := 0; j < 32; j++ {
		for i := 0; i < 8; i++ {
			if c1[j]<<i != c2[j]<<i {
				blorg++
			}
		}
	}

	fmt.Printf("Diffcount: %d\n", blorg)
}
