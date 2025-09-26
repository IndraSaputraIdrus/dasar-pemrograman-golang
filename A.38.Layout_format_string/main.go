package main

import "fmt"

type Student struct {
	name        string
	height      float32
	age         int8
	isGraduated bool
	hobbies     []string
}

func main() {
	data := Student{
		name:        "indra",
		height:      165,
		age:         22,
		isGraduated: true,
		hobbies:     []string{"gaming", "coding", "read novel", "read manhwa"},
	}

	// %b untuk biner
	fmt.Printf("%b \n", data.age)

	// %c untuk data numberic ke bentuk string unicode
	fmt.Printf("%c \n", 1235)

	// %d untuk number ( int )
	fmt.Printf("%d \n", data.age)

	// %e dan %e untuk untuk memformat data numerik desimal ke dalam bentuk notasi numerik standar Scientific notation.
	fmt.Printf("%e \n", data.height)
	fmt.Printf("%E \n", data.height)

	// %f dan %F (kedua nya sama) untuk format numberik desimal (default 6 digit)
	fmt.Printf("%f \n", data.height)
	fmt.Printf("%.9f \n", data.height)
	fmt.Printf("%.2f \n", data.height)
	fmt.Printf("%.f \n", data.height)

	// %G adalah alias dari %g. Keduanya memiliki fungsi yang sama.
	// untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan.
	fmt.Printf("%g \n", 0.123123123123)

	// %o untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan.
	fmt.Printf("%o \n", data.age)

	// %p untuk format data pointer
	fmt.Printf("%p \n", &data.name)

	// %q untuk espace string
	fmt.Printf("%q \n", `" name \ height "`)

	// %s untuk format string
	fmt.Printf("%s \n", data.name)

	// %t untuk format boolean
	fmt.Printf("%t \n", data.isGraduated)

	// %T untuk ambil type
	fmt.Printf("%T \n", data.isGraduated)
	fmt.Printf("%T \n", data.name)
	fmt.Printf("%T \n", data.height)

	// %v untuk semua termasuk 'any'
	fmt.Printf("%v \n", data.name)
	fmt.Printf("%v \n", data.hobbies)
	fmt.Printf("%v \n", data.isGraduated)
	fmt.Printf("%v \n", data)

	// %+v untuk format struct
	fmt.Printf("%+v \n", data)

	// %#v untuk memformat struct, 
	// mengembalikan nama dan nilai tiap property sesuai dengan struktur struct dan 
	// juga bagaimana objek tersebut dideklarasikan
	fmt.Printf("%#v \n", data)

	// %x atau %X untuk format string numerik heksadesimal
	fmt.Printf("%x \n", data.age)
	fmt.Printf("%x \n", data.name)

	// %% untuk menulis karakter % pada string format.
	fmt.Printf("%% \n")
}
