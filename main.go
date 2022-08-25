package main

import (
	"fmt"
)

var saldo int = 500000
var menu string
var totalSaldo, sisaSaldo, setoranSaldo, tarikSaldo int

func main() {

	kode := "12"
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
			fmt.Println("1. Menu Infromasi Saldo")
			fmt.Printf("SALDO REKENING ANDA: %b \n", totalSaldo)
		case "2":
			fmt.Println("2. Menu Setor Tunai")
			fmt.Printf("Masukkan nominal: %b", setoranSaldo)
			fmt.Scan(&setoranSaldo)
			operatorSetoran(totalSaldo, setoranSaldo)
			// fmt.Print(totalSaldo)
			// Template()
		case "3":
			fmt.Println("3. Menu Tarik Tunai")
			fmt.Printf("Masukkan nominal setoran: %b \n", tarikSaldo)
			fmt.Scan(&tarikSaldo)
		default:
			fmt.Println("Tidak terdapat menu")
			Template()
		}
	}
}

func Template() {
	fmt.Printf("SALDO ANDA: %d \n", saldo)
	fmt.Println("======= Silahkan Pilih Menu =======")
	fmt.Println("1. Informasi Saldo")
	fmt.Println("2. Setor Tunai")
	fmt.Println("3. Tarik Tunai")
	fmt.Println("4. Ubah Pin")
	fmt.Print("Silahkan Masukkan menu: ")
	fmt.Scan(&menu)
	// fmt.Print("======= Copyright by Dafrin =======")
}

func operatorSetoran(totalSaldo, setoranSaldo int) int {
	var value int = totalSaldo + setoranSaldo
	fmt.Printf("Saldo anda: %b", value)
	return value
}
