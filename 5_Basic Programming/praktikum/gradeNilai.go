package main

import "fmt"

func main() {
	// input
	var nilaiSiswa int
	var namaSiswa string


	fmt.Print("Masukkan Nama Siswa : ")
	fmt.Scanln(&namaSiswa)
	fmt.Print("Masukkan Nilai Siswa : ")
	fmt.Scanln(&nilaiSiswa)

	// Precess: Your Solution Code Here
	fmt.Print(namaSiswa, " mendapatkan ")
	if nilaiSiswa < 0 && nilaiSiswa > 100 {
		fmt.Println("Nilai Invalid")
	} else if nilaiSiswa >= 80 {
		fmt.Println("Nilai A")
	} else if nilaiSiswa >= 65 && nilaiSiswa <= 79 {
		fmt.Println("Nilai B")
	} else if nilaiSiswa >= 50 && nilaiSiswa <= 64 {
		fmt.Println("Nilai C")
	} else if nilaiSiswa >= 35 && nilaiSiswa <= 49 {
		fmt.Println("Nilai D")
	} else if nilaiSiswa <= 34 {
		fmt.Println("Nilai E")
	}


	// cara 2 

	// switch  {
	// case nilaiSiswa >= 80:
	// 	fmt.Println("A")
	// case nilaiSiswa >= 65:
	// 	fmt.Println("B")
	// case nilaiSiswa >= 50:
	// 	fmt.Println("C")
	// case nilaiSiswa >= 35:
	// 	fmt.Println("D")
	// case nilaiSiswa >= 34:
	// 	fmt.Println("E")

	// }
	
}