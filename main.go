package main

import "fmt"

func main() {
	var nama string = "Jerry"
	var cewekBenar bool = true
	var umur int = 20
	var piLingkaran float32 = 3.14

	var subyek string // ""
	var angka int     // 0
	var koma float32  // 0
	var check bool    // false

	fmt.Println(nama, cewekBenar, umur, piLingkaran)
	fmt.Println(subyek, angka, koma, check)
	fmt.Println("====================")
	var angka1 int = 5
	var angka2 int = 2
	fmt.Println("Pembagian menggunakan int ", angka1/angka2)
	var angka3 float32 = 4
	var angka4 float32 = 2
	fmt.Println("Pembagian menggunakan float ", angka3/angka4)

	hasil1 := angka1 / angka2
	hasil1 = int(angka3) / int(angka4)
	var hasil2 float32
	hasil2 = float32(angka1) / float32(angka2)
	hasil3 := angka3 + float32(angka1)
	fmt.Println("Hasil 1", hasil1)
	fmt.Println("Hasil 2", hasil2)
	fmt.Println("Hasil 3", hasil3)

	fullName := "Jerry Young"
	phon_number := "08123456789"
	fmt.Println(fullName, phon_number)

	fmt.Println(fullName, phon_number)

	const pi float32 = 3.14
	fmt.Println(pi)

}
