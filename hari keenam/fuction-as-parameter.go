package main

import "fmt"

func sayHelloWithFillter(name string, filter func(string) string) {
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
