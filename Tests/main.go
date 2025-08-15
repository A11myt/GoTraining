package main

import (
	"fmt"
)

func main() {
	// operationTable()
	// logicShift()
	overflow()
}

func operationTable() {

	var a uint = 0b00001010
	var b uint = 0b00001100

	fmt.Printf("    a: %08b\n", a)
	fmt.Printf("not a: %08b\n", ^a)
	fmt.Println()
	fmt.Printf("a AND b: %08b\n)", a&b)
	fmt.Printf("a OR b: %08b\n", a|b)
	fmt.Printf("a XOR b: %08b\n", a^b)
	fmt.Printf("a AND NOT b: %08b\n", a&^b)
}

func logicShift() {
	var a uint8 = 13
	var aLeft uint8 = a << 1
	var aRight uint8 = a >> 1
	fmt.Printf("a: Bin: %08b Dez: %3v\n", a, a)
	fmt.Printf("a << 1: Bin: %08b Dez: %3v\n", aLeft, aLeft)
	fmt.Printf("a >> 1: Bin: %08b Dez: %3v\n", aRight, aRight)

	var c int8 = -13
	var cLeft int8 = c << 1
	var cRight int8 = c >> 1
	fmt.Printf("c: Bin: %08b Dez: %3v\n", c, c)
	fmt.Printf("c << 1: Bin: %08b Dez: %3v\n", cLeft, cLeft)
	fmt.Printf("c >> 1: Bin: %08b Dez: %3v\n", cRight, cRight)
}

func overflow() {
	var f uint = 200
	fmt.Println(f + 80)
	f = f + 80
	f = 200 + 80
}
