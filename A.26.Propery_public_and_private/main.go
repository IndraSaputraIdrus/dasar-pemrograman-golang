package main

import (
	"fmt"
	"properti_public_dan_private/library"

	// Import Dengan Prefix Tanda Titik
	. "properti_public_dan_private/library2"

	// alias
	t "time"
)

func main() {
	library.SayHello()

	myPerson := library.Person{
		Name:    "Indra",
		Age:     22,
		Address: "Jl. sukamaju 2",
	}

	fmt.Println("Name:", myPerson.Name)
	fmt.Println("Age:", myPerson.Age)
	fmt.Println("Address:", myPerson.Address)

	// Seperti yang kita tahu, untuk mengakses fungsi/struct/variabel yg berada di package lain,
	// nama package nya perlu ditulis,
	// contohnya seperti pada penggunaan library.Student dan fmt.Println().


    // PERINGATAN!
    //
    // Penggunaan tanda titik pada saat import package bisa menyebabkan kode menjadi ambigu, 
	// karena alasan tersebut teknik import ini kurang direkomendasikan.

	// tidak perlu library2.Sum
	total := Sum([]int{20, 1, 2, 5})
	fmt.Println(total)

	fmt.Println(t.DateTime)
}
