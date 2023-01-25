package main

import "fmt"

// create own type, underlying type be in `int`
// source: https://go.dev/ref/spec#Types
type OwnType int

func main() {
	var x OwnType
	fmt.Println(x)
	// print type of variable
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Printf("value: %v type: %T\n", x, x)

	var y = int(x)
	s := fmt.Sprintf("value: %v type: %T", y, y)
	fmt.Println(s)

}
