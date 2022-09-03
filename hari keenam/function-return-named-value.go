package main

import "fmt"

func fullName() (fristName, lastName string) {
	fristName = "fuad"
	lastName = "maulana"
	return
}

func main() {
	firstName, lastName := fullName()

	fmt.Println(firstName)
	fmt.Println(lastName)
}
