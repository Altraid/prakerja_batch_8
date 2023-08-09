package main

import "fmt"

// Fungsi untuk mengecek apakah sebuah bilangan adalah kelipatan 7 atau tidak
func isKelipatan7(n int) bool {
	// Jika n habis dibagi 7, maka kelipatan 7
	if n%7 == 0 {
		return true
	}
	// Jika tidak, maka bukan kelipatan 7
	return false
}

func main() {
	// Buat variabel n dan beri nilai awal 0
	var n int
	n = 0

	// Cetak pesan "Masukkan sebuah bilangan: "
	fmt.Print("Masukkan sebuah bilangan: ")

	// Baca nilai input dari pengguna dan simpan dalam variabel n
	fmt.Scanln(&n)

	// Panggil fungsi isKelipatan7 dan simpan hasilnya dalam variabel k
	k := isKelipatan7(n)

	// Cek nilai k
	if k {
		// Jika k bernilai true, cetak pesan "Bilangan %d adalah kelipatan 7"
		fmt.Printf("Bilangan %d adalah kelipatan 7\n", n)
	} else {
		// Jika k bernilai false, cetak pesan "Bilangan %d bukan kelipatan 7"
		fmt.Printf("Bilangan %d bukan kelipatan 7\n", n)
	}
}
