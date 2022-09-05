package main

import "fmt"

type Hasname interface {
	getName() string
}

func sayHello(hasname Hasname) {
	fmt.Println("halo", hasname.getName())
}

type Orang struct {
	name string
}

func (orang Orang) getName() string {
	return orang.name
}

func main() {
	var fuad Orang
	fuad.name = "joe"

	sayHello(fuad)
}
