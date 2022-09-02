package main

import "fmt"

func main() {
	var (
		name1 = "fuad"
		name2 = "fuad"
	)

	result := name1 == name2

	fmt.Println(result)

	var (
		umur1 = 1
		umur2 = 2
	)

	fmt.Println(umur1 > umur2)
	fmt.Println(umur1 < umur2)
	fmt.Println(umur1 >= umur2)
	fmt.Println(umur1 != umur2)
	fmt.Println(umur1 == umur2)

}
