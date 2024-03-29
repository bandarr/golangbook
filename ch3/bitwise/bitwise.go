package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) //"00100010", the set {1,5}
	fmt.Printf("%08b\n", y) //"00000110", the set {1,2}

	fmt.Printf("%08b\n", x&y)  //"00000010", the intersetion {1}
	fmt.Printf("%08b\n", x|y)  //"00100110", the union {1,2,5}
	fmt.Printf("%08b\n", x^y)  //"00100100", the symmetri differene {2,5}
	fmt.Printf("%08b\n", x&^y) //"00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { //membership test
			fmt.Println(i) //"1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) //"01000100", the set {2,6}
	fmt.Printf("%08b\n", x>>1) //"00010001", the set {0,4}

	o := 0666
	fmt.Printf("%d  %[1]o   %#[1]o\n", o)

	yerp := int64(0xdeadbeef)
	fmt.Printf("%d  %[1]x %#[1]x  %#[1]X\n", yerp)

	ascii := 'a'
	unicode := 'כ'

	newline := '\n'

	fmt.Printf("%d  %[1]c  %[1]q\n", ascii)
	fmt.Printf("%d  %[1]c  %[1]q\n", unicode)
	fmt.Printf("%d  %[1]q\n", newline)

}
