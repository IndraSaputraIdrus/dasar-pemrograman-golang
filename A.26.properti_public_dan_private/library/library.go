package library

import "fmt"

func init() {
	fmt.Println("library/library imported")
}

// Fungsi SayHello(), level aksesnya adalah publik, ditandai dengan nama fungsi diawali huruf besar.
func SayHello() {
	fmt.Println("Hello world!")
	introduce("indra")
}

// Fungsi introduce() dengan level akses private, ditandai oleh huruf kecil di awal nama fungsi.
func introduce(name string) {
	fmt.Println("Hello my name is", name)
}
