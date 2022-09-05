package main

import "fmt"

type Orang struct {
	Name, Address string
	Age           int
}

func main() {
	var fuad Orang
	fuad.Name = "teuku Fuad maulana"
	fuad.Address = "Lhokseumawe"
	fuad.Age = 30

	fmt.Println(fuad.Name)
	fmt.Println(fuad.Address)
	fmt.Println(fuad.Age)

	maulana := Orang{
		Name:    "Maulana",
		Address: "Lhokseumawe",
		Age:     30,
	}

	fmt.Println(maulana)
}
