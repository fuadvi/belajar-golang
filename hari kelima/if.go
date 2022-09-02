package main

import "fmt"

func main() {
	name := "fuad"

	if name == "fuad" {
		fmt.Println("halo " + name)
	} else if name == "maulana" {
		fmt.Println("halo " + name)
	} else if name == "teuku" {
		fmt.Println("halo " + name)
	} else {
		fmt.Println("hi, boleh kenalan ? ")
	}

	// short stament if

	if lenght := len(name); lenght > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
