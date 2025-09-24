// Reflection adalah teknik untuk inspeksi variabel,
// mengambil informasi dari suatu variabel untuk dilihat metadatanya atau untuk keperluan manipulasi.
// Cakupan informasi yang bisa didapatkan lewat reflection sangat luas,
// seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.

// Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe reflect.Value,
// yang berisikan informasi yang berhubungan dengan nilai/data variabel yang diinspeksi.

// Sedangkan reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type.
// Objek tersebut berisikan informasi yang berhubungan dengan tipe data variabel yang diinspeksi.

package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	rv := reflect.ValueOf(s)

	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}

	rvType := rv.Type()

	for v := range rv.NumField() {
		fmt.Println("nama:", rvType.Field(v).Name)
		fmt.Println("tipe data:", rvType.Field(v).Type)
		fmt.Println("nilai:", rv.Field(v).Interface())
		fmt.Println("---")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	number := 23
	reflectValue := reflect.ValueOf(number)

	fmt.Println("type variabel:", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable:", reflectValue.Int())
	}

	nilai := reflectValue.Interface().(int)
	fmt.Println(nilai)

	indra := &student{
		Name:  "Indra",
		Grade: 2,
	}

	indra.getPropertyInfo()

	fmt.Println("my name:", indra.Name)
	rv := reflect.ValueOf(indra)
	method := rv.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("indra saputra idrus"),
	})
	fmt.Println("my new name:", indra.Name)
}
