package main

import "fmt"

func main() {
	var p, q, r int
	var s bool

	fmt.Print("Masukkan kode digit: ")
	fmt.Scan(&p)

	q = p / 1000
	r = p % 10

	s = q == r

	fmt.Print(s)
}