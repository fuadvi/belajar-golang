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

	// short hand switch expresseion
	// switch lenght := len(name); lenght > 5 {
	// case true:
	// 	fmt.Println("nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	// switch tanpa kondisi
	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("nama terlalu panjang")
	case lenght > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah sesuai")
	}

}
