package main

import (
	"fmt"
)

// Struktur data untuk barang
type Barang struct {
	Nama  string
	Harga float64
	Stok  int
}

func main() {
	// Membuat daftar barang
	daftarBarang := []Barang{
		{"Rokok Surya", 25.0, 70},
		{"Rokok Sampoerna", 30.0, 80},
		{"Rokok Diplomat", 26.0, 60},
		{"Rokok WinBold", 27.0, 75},
	}

	// Menampilkan daftar barang
	fmt.Println("Daftar Barang yang Tersedia:")
	for i, barang := range daftarBarang {
		fmt.Printf("%d. %s (Harga: %.2f, Stok: %d)\n", i+1, barang.Nama, barang.Harga, barang.Stok)
	}

	// Inisialisasi variabel total dan keranjang belanja
	var total float64
	keranjangBelanja := make(map[string]int)

	// Inputan pengguna
	for {
		var pilihan int
		var jumlah int

		fmt.Print("Pilih barang (nomor) atau ketik 0 untuk selesai belanja: ")
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			break
		}

		if pilihan < 1 || pilihan > len(daftarBarang) {
			fmt.Println("Pilihan tidak valid.")
			continue
		}

		fmt.Print("Jumlah barang yang dibeli: ")
		fmt.Scan(&jumlah)

		barang := daftarBarang[pilihan-1]

		if jumlah > barang.Stok {
			fmt.Println("Stok tidak mencukupi.")
			continue
		}

		// Menambahkan barang ke keranjang belanja
		keranjangBelanja[barang.Nama] += jumlah

		// Mengurangi stok barang
		daftarBarang[pilihan-1].Stok -= jumlah

		// Menghitung total
		total += barang.Harga * float64(jumlah)
	}

	// Menampilkan total dan cetak hasil total
	fmt.Printf("Total biaya: %.2f\n", total)

	fmt.Println("Rincian Belanja:")
	for barang, jumlah := range keranjangBelanja {
		fmt.Printf("%s: %d\n", barang, jumlah)
	}
}
