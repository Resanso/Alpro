package main

import "fmt" 

func main() {
	var x, N int
	var v bool

	fmt.Println("masukkan bilangan bulat x dan N: ")
	fmt.Scan(&x, &N)	
	v = x % N == 0

	fmt.Printf("%d kelipatan %d? %t", x, N, v)
}
