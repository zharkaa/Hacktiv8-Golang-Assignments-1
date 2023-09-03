package main

import (
	"fmt"
	"os"
)

type Friends struct {
	ID        int
	nama      string
	umur      int
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error: Nama harus diisi")
		fmt.Println("Example: main.go Aurora")
		return
	}

	a := os.Args[1]

	person := []Friends{
		{1, "Aurora", 20, "Manado", "Programmer", "Mencari pengalaman dan menambah ilmu"},
		{2, "Alvaro", 22, "Berlin", "Mahasiswa", "Menambah skill tentang Golang"},
		{3, "Invoker", 19, "Tokyo", "Mahasiswa", "Hacking NASA"},
	}

	found := false
	
	for _, v := range person {
		if a == v.nama {
			printFriends(v)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Nama tidak ditemukan")
	}
}

func printFriends(f Friends) {
	fmt.Println("Nama :", f.nama)
	fmt.Println("Umur :", f.umur)
	fmt.Println("Alamat :", f.alamat)
	fmt.Println("Pekerjaan :", f.pekerjaan)
	fmt.Println("Alasan :", f.alasan)
}
