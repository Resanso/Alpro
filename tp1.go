package main

import "fmt"

func main() {
	var p int
	var q bool

	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&p)

	q = p % 4 == 0 && p % 5 == 0 && p % 100 != 0

	fmt.Print(q)
}