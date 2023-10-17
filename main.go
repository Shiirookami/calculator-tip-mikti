package main

import "fmt"

func main() {
	// Data uji
	tagihan1 := 275.0
	tagihan2 := 40.0
	tagihan3 := 432.0

	// menghitung tip
	hitungTip := func(tagihan float64) float64 {
		if tagihan >= 50 && tagihan <= 300 {
			return tagihan * 0.15
		} else {
			return tagihan * 0.2
		}
	}

	// Menghitung tip dan total untuk setiap data uji
	tip1 := hitungTip(tagihan1)
	tip2 := hitungTip(tagihan2)
	tip3 := hitungTip(tagihan3)

	total1 := tagihan1 + tip1
	total2 := tagihan2 + tip2
	total3 := tagihan3 + tip3

	// Menampilkan hasil
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan1, tip1, total1)
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan2, tip2, total2)
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan3, tip3, total3)
}
