package main

import "fmt"

func main() {
	var jumlah int
	fmt.Scan(&jumlah)

	// loop dimulai dari 0 berjalan hingga sebanyak jumlah
	// setiap iterasi berarrti satu perhitungan luas segitiga
	for i := 0; i < jumlah; i++ {
		// deklarasi untuk alas dan tinggi
		var alas, tinggi float64
		fmt.Scan(&alas, &tinggi)

		// rumus luas segitiga 1/2 * a * t
		luas := 0.5 * alas * tinggi
		// %.0f = menampilkan nilai alas dan tinggi sebagai bilangan bulat bukan desimal
		// %.1f = menampilkan hasil perhitungan luas dengan satu angka di belakang koma
		fmt.Printf("(%.0f * %.0f) / 2 = %.1f\n", alas, tinggi, luas)
	}
}
