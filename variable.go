package main

import "fmt"

func main() {
	var name string
	var age = 10
	lastName := "Suchi Hafizullah"

	var (
		address  = "Muara Enim"
		province = "South Sumatera"
	)

	fmt.Println(age)

	name = "Adam Suchi Hafizullah"
	fmt.Println(name)

	name = "Adam SH"
	age = 28
	lastName = "Hafizullah"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(lastName)

	fmt.Println(province)
	fmt.Println(address)
}
