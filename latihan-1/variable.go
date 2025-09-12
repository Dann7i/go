package main

import "fmt"

func main() {
	var nama string = "Wildan Hermawan"
	var umur int = 17
	// contoh kata kunci :=
	kota := "Kota Cimahi"
	tahun := 2025

	fmt.Println("Nama saya ", nama)
	fmt.Println("Umur saya ", umur)
	fmt.Println("Saya tinggal di ", kota)
	fmt.Println("Tahun sekarang ", tahun)

	// contoh penggunaan bervalue
	var isi string
	isi = "isi sudah diisi"
	isi = "isi sudah diisi ulang"
	fmt.Println(isi)

	// contoh penggunaan multiple variable

	var (
		a = "b"
		b = "a"
	)
	fmt.Println(a, b)

}

// var: deklarasi variable
// tipe data harus ditulis
// := : deklarasi variable tanpa menulis tipe data, tipe data akan mengikuti nilai yang diisikan
// string: teks
// int: bilangan bulat
