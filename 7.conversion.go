package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Akbar Rafsanjani"
	var e uint8 = name[6]   // ambil huruf index ke 6 dari variable name
	var eString = string(e) // konversi dari uint8 ke string

	fmt.Println(e)
	fmt.Println(eString)
}
