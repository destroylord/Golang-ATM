package main

import (
	"fmt"
)

var saldo int = 500000
var menu string

func main() {

	kode := "admin123"
	var pass string = kode

	fmt.Print("Masukkan Kode bank anda: ")
	fmt.Scan(&kode)

	if pass != kode {
		fmt.Print("Kode yang anda masukkan salah, silahkan input dengan benar \n")
		main()
	} else {
		Template()

		switch menu {
		case "1":
			fmt.Print("1. Manu Infromasi Saldo")
		case "2":
			fmt.Print("2. Manu Tambah Saldo")
		case "3":
			fmt.Print("3. Menu Tarik Tunai")
		default:
			fmt.Print("Tidak terdapat menu")
		}
	}
}

func Template() {
	fmt.Printf("SALDO ANDA: %d \n", saldo)
	fmt.Println("======= Silahkan Pilih Menu =======")
	fmt.Println("1. Informasi Saldo")
	fmt.Println("2. Tambah Saldo")
	fmt.Println("3. Tarik Tunai")
	fmt.Print("Silahkan Masukkan menu: ")
	fmt.Scan(&menu)
	// fmt.Print("======= Copyright by Dafrin =======")

}
