package main

import "fmt"

// Fungsi untuk mengecek apakah sebuah bilangan adalah prima atau tidak
func isPrima(n int) bool {
	// Jika n kurang dari 2, maka bukan prima
	if n < 2 {
		return false
	}
	// Jika n sama dengan 2, maka prima
	if n == 2 {
		return true
	}
	// Jika n habis dibagi 2, maka bukan prima
	if n%2 == 0 {
		return false
	}
	// Cek semua bilangan ganjil dari 3 sampai akar kuadrat n
	for i := 3; i*i <= n; i += 2 {
		// Jika n habis dibagi i, maka bukan prima
		if n%i == 0 {
			return false
		}
	}
	// Jika tidak ada pembagi lain, maka prima
	return true
}

func main() {
	// Variasi kode berupa memilih melanjutkan menghitung nilai atau berhenti menghitung nilai
	// Buat variabel lanjut dan beri nilai awal "y"
	var lanjut string
	lanjut = "y"

	// Buat perulangan selama lanjut bernilai "y" atau "Y"
	for lanjut == "y" || lanjut == "Y" {
		// Buat variabel n dan beri nilai awal 0
		var n int
		n = 0

		// Cetak pesan "Masukkan sebuah bilangan: "
		fmt.Print("Masukkan sebuah bilangan: ")

		// Baca nilai input dari pengguna dan simpan dalam variabel n
		fmt.Scanln(&n)

		// Panggil fungsi isPrima dan simpan hasilnya dalam variabel p
		p := isPrima(n)

		// Cek nilai p
		if p {
			// Jika p bernilai true, cetak pesan "Bilangan %d adalah prima"
			fmt.Printf("Bilangan %d adalah prima\n", n)
		} else {
			// Jika p bernilai false, cetak pesan "Bilangan %d bukan prima"
			fmt.Printf("Bilangan %d bukan prima\n", n)
		}

		// Cetak pesan "Apakah Anda ingin melanjutkan menghitung nilai? (y/n): "
		fmt.Print("Apakah Anda ingin melanjutkan menghitung nilai? (y/n): ")

		// Baca nilai input dari pengguna dan simpan dalam variabel lanjut
		fmt.Scanln(&lanjut)
	}

	// Cetak pesan "Terima kasih telah menggunakan program ini"
	fmt.Println("Terima kasih telah menggunakan program ini")
}
