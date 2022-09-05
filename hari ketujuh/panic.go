package main

import "fmt"

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("error dengan message ", message)
	}
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
	runApp(true)
}
