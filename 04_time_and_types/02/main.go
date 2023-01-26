package main

import (
	"fmt"
)

// Write a program that prints a number in decimal, binary, and hex.
func prints_values(input_from_outside int) {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	fmt.Printf("Decimal: %d\n", input)
	fmt.Printf("Binary: %b\n", input)
	fmt.Printf("Hexadecimal: %#X\n", input) // that will print the input number with a leading '0x' or '0X' depends on the case
}
