package main

import "fmt"

func main() {
	var modal float64
	fmt.Scan(&modal)

	// loop yang dimulai dari 1 hingga 10
	// setiap iterasi berarti 1 tahun
	// tahun++ berarti nilai tahun akan bertambah 1
	for tahun := 1; tahun <= 10; tahun++ {
		// meghitung nilai modal dimana modal asli 100% + 5% bunga
		// sehingga didapatkan nilai 1.05
		modal = modal * 1.05
		// %d format int
		// %.2f dua angka dibelakang koma
		fmt.Printf("Tahun ke-%d: Rp %.2f\n", tahun, modal)
	}
}
