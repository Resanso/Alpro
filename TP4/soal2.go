package main

import "fmt"

func main() {
	var tahun int
	var hasil bool

	fmt.Print("Masukkan sebuah Tahun: ")
	fmt.Scan(&tahun)

	hasil = tahun%4 == 0 && tahun%5 == 0 && tahun%100 != 0
	fmt.Println(hasil)
}
