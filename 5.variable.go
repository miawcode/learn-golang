package main

import "fmt"

func main() {
	var name string

	name = "Akbar Rafsanjani"
	fmt.Println(name)

	name = "Lorem Ipsum"
	fmt.Println(name)

	// bisa menginisialisasi variable otomatis dengan langsung menyabutkan nilainya
	var angka = 123
	fmt.Println(angka)

	barang := "raket nyamuk"
	fmt.Println(barang)

	//inisialisasi banyak variable
	var (
		firstName = "Akbar"
		lastName  = "Rafsanjani"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
