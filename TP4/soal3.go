package main

import "fmt"

func main() {
	var kode, a, b, c, d int
	var hasil bool

	fmt.Print("Masukkan 4 digit kode: ")
	fmt.Scan(&kode)

	a = kode / 1000
	b = (kode / 1000) % 10
	c = kode % 100
	d = kode % 10

	hasil = a == d && b != c
	fmt.Println(hasil)

}
