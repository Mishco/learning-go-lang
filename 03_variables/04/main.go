package main

import "fmt"

// create own type, underlying type be in `int`
// source: https://go.dev/ref/spec#Types
type OwnType int

func main() {
	var x OwnType
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
