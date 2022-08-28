package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

// global variable
var menu, namaPengguna, kembali string
var totalSaldo, sisaSaldo, setoranSaldo, tarikSaldo int

func main() {

	// Converet To Currency
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2}

	// Default set Money
	var saldo string = ac.FormatMoney(1000000)

	totalSaldo = 50000
	// saldoTotal := saldo + totalSaldo
	// fmt.Println("total saldo: ", saldoTotal)
	namaPengguna = "John Doe"

	// pin
	pin := 1234
	pass := pin

	fmt.Print("Masukkan Kode bank anda: ")
	fmt.Scan(&pin)

	if pass != pin {
		fmt.Print("Kode yang anda masukkan salah, silahkan input dengan benar \n")
		main()
	} else {
		Template()

		switch menu {
		case "1":
			fmt.Println("1. Menu Infromasi Saldo")
			fmt.Printf("NamaPengguna: %s \n", namaPengguna)
			fmt.Printf("Sebelumnya: %v \n", saldo)
			fmt.Print("<- Kembali")
			// fmt.Printf("SALDO REKENING ANDA: Rp.%d \n", saldoTotal)

		case "2":
			fmt.Printf("Saldo saat ini: %v \n", totalSaldo)
			fmt.Println("2. Menu Setor Tunai")
			fmt.Printf("Masukkan nominal: ")
			fmt.Scan(&setoranSaldo)
			operatorSetoran(totalSaldo, setoranSaldo)
			// fmt.Print(totalSaldo)
			// Template()
		case "3":
			fmt.Println("3. Menu Tarik Tunai")
			// fmt.Println("Saldo Awal : ", saldoTotal)
			// OperatorTarikTunai(saldoTotal)
			// fmt.Printf("Masukkan nominal setoran: %b \n", tarikSaldo)
			// fmt.Scan(&tarikSaldo)
		default:
			fmt.Println("Tidak terdapat menu")
			Template()
		}
	}
}

func Template() {
	// fmt.Printf("SALDO ANDA: %d \n", saldo)
	fmt.Println("======= Silahkan Pilih Menu =======")
	fmt.Println("1. Informasi Saldo")
	fmt.Println("2. Setor Tunai")
	fmt.Println("3. Tarik Tunai")
	fmt.Println("4. Ubah Pin")
	fmt.Print("Silahkan Masukkan menu: ")
	fmt.Scan(&menu)
	fmt.Println("\033[2J")
	// fmt.Print("======= Copyright by Dafrin =======")
}

func operatorSetoran(totalSaldo, setoranSaldo int) int {
	var value int = totalSaldo + setoranSaldo
	fmt.Printf("Saldo anda: %b", value)
	return value
}

func OperatorTarikTunai(totalSaldo int) int {
	var tarikSaldo int
	minimumSaldo := 6000
	fmt.Print("jumlah tarik saldo: ")
	fmt.Scan(&tarikSaldo)

	outputSaldo := totalSaldo - tarikSaldo
	if totalSaldo > minimumSaldo {
		fmt.Printf("maaf saldo anda tidak boleh kurang dari %d \n", minimumSaldo)
	} else if totalSaldo < minimumSaldo {
		fmt.Printf("saldo saat ini : %d", outputSaldo)
	}
	return outputSaldo

}
