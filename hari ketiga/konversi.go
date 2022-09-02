package main

import "fmt"

func main() {

	// konversi tipe data integer
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi byte to string
	name := "fuad"
	e := name[0]
	converstionString := string(e)

	fmt.Println(name)
	fmt.Println(converstionString)
}
