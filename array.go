package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Adam"
	names[1] = "Suchi"
	names[2] = "Hafizullah"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90, 29, 39,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))

	values[0] = 100
	fmt.Println(values[0])
}
