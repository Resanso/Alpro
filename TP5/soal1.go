package main

import "fmt"

func main() {
	// deklarasi variabel
	var terbesar, terkecil, hasil int
	fmt.Scan(&terbesar, &terkecil)

	// loop dimulai dari nilai terkecil hingga terbesar
	// i++ menambahkan 1 setiap iterasi
	for i := terkecil; i <= terbesar; i++ {
		fmt.Print(i)
		if i < terbesar {
			// untuk memberikan spasi setelah setiap angka
			fmt.Print(" ")
		}
		// menambahkan nilai i ke dalam variabel hasil
		// untuk mendapatkan jumlah total dari seluruh angka
		hasil += i
	}
	fmt.Println()
	fmt.Print("Total = ", hasil)
}
