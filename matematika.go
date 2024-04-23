package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	var i = 10
	i += 10

	var j = 1
	j++

	fmt.Println(c)
	fmt.Println(i)
	fmt.Println(j)

	j--
	fmt.Println(j)
}
