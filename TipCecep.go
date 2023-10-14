package main

import (
	"fmt"
)

func main() {
	// Ubah nilai tagihan sesuai dengan yang diinginkan
	tagihan1 := 275
	tagihan2 := 40
	tagihan3 := 430

	// Fungsi untuk menghitung tip berdasarkan nilai tagihan
	calculateTip := func(tagihan int) (float64, float64) {
		var tip, total float64

		if tagihan >= 50 && tagihan <= 300 {
			tip = float64(tagihan) * 0.15 // 15% tip
		} else {
			tip = float64(tagihan) * 0.20 // 20% tip
		}

		total = float64(tagihan) + tip
		return tip, total
	}

	// Hitung tip dan tampilkan hasil untuk data uji
	tip1, total1 := calculateTip(tagihan1)
	tip2, total2 := calculateTip(tagihan2)
	tip3, total3 := calculateTip(tagihan3)

	// Tampilkan hasil di konsol
	fmt.Printf("Tagihannya %d, tipnya %.2f, dan total nilainya %.2f\n", tagihan1, tip1, total1)
	fmt.Printf("Tagihannya %d, tipnya %.2f, dan total nilainya %.2f\n", tagihan2, tip2, total2)
	fmt.Printf("Tagihannya %d, tipnya %.2f, dan total nilainya %.2f\n", tagihan3, tip3, total3)
}
