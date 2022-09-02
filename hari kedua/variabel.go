package main

import "fmt"

func main() {
	// membuat variabel dengan var dan inisialisasis setelah nya
	var nama string
	nama = "TEUKU FUAD MAULANA"

	fmt.Println(nama)

	nama = "ika wulandari"
	fmt.Println(nama)

	// membuat variabel dengan var dan langsung inisialisasis
	var freindName = "teuku fuad maulana"
	fmt.Println(freindName)

	// membuat variabel tanpa mengunakan var
	//  var tidak wajib digunakan bisa di ganti menjadi :=
	country := "indonesia"
	fmt.Println(country)

	// multipel variabel
	var (
		fristName = "Teuku"
		lastName  = "fuad maulana"
	)

	fmt.Println(fristName)
	fmt.Println(lastName)
}
