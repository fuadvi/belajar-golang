package main

import "fmt"

type Blacklist func(string) bool

func userRegister(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked")
	} else {
		fmt.Println("welcom " + name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	userRegister("admin", blacklist)
	userRegister("user", blacklist)
}
