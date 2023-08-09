package main

import "fmt"

// Fungsi untuk menghitung luas trapesium
func luasTrapesium(a, b, t float64) float64 {
	// Rumus luas trapesium adalah 1/2 * (a + b) * t
	return 0.5 * (a + b) * t
}

func main() {
	// Variasi kode berupa nilai input a, b, t manual
	// Buat variabel a, b, t dan beri nilai awal 0
	var a, b, t float64
	a = 0
	b = 0
	t = 0

	// Cetak pesan "Masukkan nilai sisi atas trapesium (cm): "
	fmt.Print("Masukkan nilai sisi atas trapesium (cm): ")

	// Baca nilai input dari pengguna dan simpan dalam variabel a
	fmt.Scanln(&a)

	// Cetak pesan "Masukkan nilai sisi bawah trapesium (cm): "
	fmt.Print("Masukkan nilai sisi bawah trapesium (cm): ")

	// Baca nilai input dari pengguna dan simpan dalam variabel b
	fmt.Scanln(&b)

	// Cetak pesan "Masukkan nilai tinggi trapesium (cm): "
	fmt.Print("Masukkan nilai tinggi trapesium (cm): ")

	// Baca nilai input dari pengguna dan simpan dalam variabel t
	fmt.Scanln(&t)

	// Panggil fungsi luasTrapesium dan simpan hasilnya dalam variabel l
	l := luasTrapesium(a, b, t)

	// Cetak hasilnya dengan format "Luas trapesium dengan sisi atas %.2f cm, sisi bawah %.2f cm, dan tinggi %.2f cm adalah %.2f cm^2"
	fmt.Printf("Luas trapesium dengan sisi atas %.2f cm, sisi bawah %.2f cm, dan tinggi %.2f cm adalah %.2f cm^2\n", a, b, t, l)
}
