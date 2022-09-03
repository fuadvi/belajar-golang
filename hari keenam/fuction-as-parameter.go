package main

import "fmt"

type Filter func(string) string

func sayHelloWithFillter(name string, filter Filter) {
	filterNamed := filter(name)
	fmt.Println("hallo " + filterNamed)
}

func spamFillter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

func main() {
	sayHelloWithFillter("fuad", spamFillter)
	sayHelloWithFillter("Anjing", spamFillter)

}
