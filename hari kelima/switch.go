package main

import "fmt"

func main() {
	name := "fuaddd"

	switch name {
	case "fuad":
		fmt.Println("hi " + name)

	case "maulana":
		fmt.Println("Hi ", name)

	default:
		fmt.Println("Hi boleh kenalan ?")
	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("nama Terlalu Panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

}
