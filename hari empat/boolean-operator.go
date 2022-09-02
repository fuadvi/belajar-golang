package main

import "fmt"

func main() {
	var ujian = 81
	var absensi = 81

	var lulusUjian = ujian > 80
	var lulusAbsensi = absensi > 80

	var lulus = lulusAbsensi && lulusUjian

	fmt.Println(lulus)

}
