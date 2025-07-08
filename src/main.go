package main

import (
	"fmt"
	"successor/components"
)

func main() {
	var a components.NaturalNumber = components.Int(563124)
	var b components.NaturalNumber = components.Int(123589)

	fmt.Println(components.Monus(a, b))
}
