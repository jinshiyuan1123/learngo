package main

import (
	"fmt"
	"math/cmplx"
)

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func main() {
	euler()
}
