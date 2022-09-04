package main

import "fmt"

func endApp() {
	fmt.Println("aplikasi selesai berjalan")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}

	fmt.Println("apliakasi jalan")
}

func main() {
	runApp(false)
}
