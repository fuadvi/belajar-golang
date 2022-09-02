package main

import (
	"fmt"
)

func main() {
	var person map[string]string = map[string]string{
		"name":    "fuad",
		"address": "lhokseumawe",
	}

	// menampah value baru map
	person["sex"] = "male"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["sex"])

	// membuat map tanpa ada data

	materi := make(map[string]string)

	materi["judul"] = "membuat bracket tournamen"
	materi["upss"] = "salah"

	fmt.Println(materi)

	// hapus value map
	delete(materi, "upss")
	fmt.Println(materi)
}
