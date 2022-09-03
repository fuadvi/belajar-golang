package main

import "fmt"

func fullName() (string, string) {
	return "fuad", "maulana"
}

func main() {
	firstName, lastName := fullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
