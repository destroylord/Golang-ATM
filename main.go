package main

import (
	"fmt"
)

// global variable
var menu, namaPengguna, kembali string
var totalSaldo, sisaSaldo, setoranSaldo, tarikSaldo int

func main() {

	// Default set Money
	saldo := 1000000

	totalSaldo = 50000
	// saldoTotal := saldo + totalSaldo
	// fmt.Println("total saldo: ", saldoTotal)
	namaPengguna = "John Doe"

	var namaPengguna string = "asdasd"

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

		case "2":
			fmt.Printf("Saldo saat ini: %v \n", totalSaldo)
			fmt.Println("2. Menu Setor Tunai")
			fmt.Print("Kirim Nominal :")
			var nominalSetor int
			fmt.Scan(&nominalSetor)
			totalSetoranSaldo := totalSaldo + nominalSetor
			fmt.Printf("Setoran kamu telah ditambahkan, saldo saat ini: %d", totalSetoranSaldo)

		case "3":
			fmt.Println("3. Menu Tarik Tunai \n")
			fmt.Println("1. 50000")
			fmt.Println("2. 100000")
			fmt.Println("3. 200000")
			fmt.Println("4. 500000")
			fmt.Println("5. 1000000")
			fmt.Println("6. Jumlah lainnya ")

			// dibikin array kemudian di looping
			var menus = map[int]int{
				1: 50000,
				2: 100000,
				3: 200000,
				4: 500000,
				5: 1000000,
			}

			for key := range menus {

				fmt.Printf("Silahkan masukkan menu: ")
				fmt.Scan(&key)

				if key >= len(menus) {

					var changeNominal int
					fmt.Printf("Masukkan nominal yang anda inginkan: ")
					fmt.Scan(&changeNominal)

					if changeNominal <= 50000 {
						fmt.Printf("Maaf minimum penarikan saldo sebesar Rp.50.000,00")
						break
					} else {

						totalPenarikanSaldoChange := (saldo - changeNominal) - 2500

						fmt.Printf("Sisa saldo anda setelah penarikan adalah :%d", totalPenarikanSaldoChange)
						break
					}

				} else {

					totalPenarikanSaldo := (saldo - menus[key]) - 2500
					fmt.Printf("Saldo penarikan saat ini: %d", totalPenarikanSaldo)
					break
				}

			}

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
	// fmt.Print("======= Copyright by Dafrin =======")
}

// func OperatorTarikTunai(totalSaldo int) int {
// 	var tarikSaldo int
// 	minimumSaldo := 6000
// 	fmt.Print("jumlah tarik saldo: ")
// 	fmt.Scan(&tarikSaldo)

// 	outputSaldo := totalSaldo - tarikSaldo
// 	if totalSaldo > minimumSaldo {
// 		fmt.Printf("maaf saldo anda tidak boleh kurang dari %d \n", minimumSaldo)
// 	} else if totalSaldo < minimumSaldo {
// 		fmt.Printf("saldo saat ini : %d", outputSaldo)
// 	}
// 	return outputSaldo

// }
