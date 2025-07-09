package main

import (
	"fmt"
	"successor/components"
)

func main() {
	var a components.NaturalNumber = components.Int(552978)
	var b components.NaturalNumber = components.Int(4)

	fmt.Println(components.Exponents(a, b))
}
