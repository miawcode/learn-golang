package main

import (
	"fmt"
)

func main() {
	const phi = 3.14

	fmt.Println("Nilai phi adalah", phi)

	// const apabila tidak dipakai, tidak terjadi error
	// berbeda dengan variable biasa, error jika tidak dipakai

	const (
		firstName = "Akbar"
		lastName  = "Rafsanjani"
	)
}
