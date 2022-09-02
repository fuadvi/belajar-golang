package main

import "fmt"

func main() {
	type NoKtp string
	var ektp NoKtp = "125215132"
	fmt.Println(ektp)
}
