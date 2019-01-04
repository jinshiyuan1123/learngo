package main

import "fmt"

func variableZeroValue() {
	var a, b int = 3, 4
	var s string = "php菜鸟书院"
	fmt.Println(a, b, s)
}

func variableShort() {
	a, b, c, s := 2, 4, true, "default"
	b = 5
	fmt.Println(a, b, c, s)
}
func main() {
	fmt.Println("hello")
	variableZeroValue()
	variableShort()
}
