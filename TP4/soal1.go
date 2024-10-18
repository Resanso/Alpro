package main

import "fmt"

func main() {
	var bilangan, kelipatan int
	var hasil bool

	fmt.Print("Inputkan sebuah bilangan dan kelipatannya: ")
	fmt.Scan(&bilangan, &kelipatan)

	hasil = bilangan%kelipatan == 0
	fmt.Println(bilangan, "merupakan kelipatan dari", kelipatan, hasil)
}
