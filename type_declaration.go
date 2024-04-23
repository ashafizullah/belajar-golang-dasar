package main

import "fmt"

func main() {
	type NoKTP string

	var ktpEko NoKTP = "22222222"
	var contoh string = "1212121"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(NoKTP("11111111"))
	fmt.Println(contohKtp)
}
