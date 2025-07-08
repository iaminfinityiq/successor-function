package main

import (
	"fmt"
	"successor/components"
)

func main() {
	var one components.NaturalNumber = components.Int(1)
	var two components.NaturalNumber = one.Add(one)

	fmt.Println(two.Add(one).Value)
}
