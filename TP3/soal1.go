package main

import "fmt"

func main(){
	var detik int
	fmt.Print("Input sebuah bilangan bulat yang bernilai positif : ")
	fmt.Scan(&detik)
	
	//Perhitungan konversi
	hitunganJam := detik / 4500
	hitunganDetik := detik / 4500
	hitunganMenit := hitunganDetik / 75
	detikMars := hitunganDetik / 75
	
	fmt.Printf("Maka hasil koversinya adalah %d jam, %d menit, dan %d detik di Mars\n", 
			    hitunganJam, hitunganMenit, detikMars)
}