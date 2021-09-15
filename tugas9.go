package main

import (
	"errors"
	"fmt"
	"strings"

)
var menus = [...]string{"Hazelnut", "Oreo", "Cappuchino", "Mocha", "Chocolate", "Butter"}
var kembali string
var data = []string{}

func validasi(input string) (bool, error) {
	var save string
	if input == "" {
		return false, errors.New("Belum ada yang dipilih.")
	}
	for i, _ := range menus {
		if strings.Title(input) == menus[i] {
			save = strings.Title(input)
		}
	}
	if save == "" {
		return false, errors.New("Tidak ada dalam daftar")
	}
	return true, nil
}



func keranjang(menus ...string) []string {
	var menu, pilih string
	for {
		fmt.Print("Masukan menu pesanan anda dalam huruf (contoh : Mocha) : ")
		fmt.Scan(&menu)
		if valid, err := validasi(menu); valid {
			data = append(data, menu)
			menu = ""
		} else {
			fmt.Println(err.Error())
		}

		fmt.Print("Ada Pesanan Lain? (y/n) : ")
		fmt.Scan(&pilih)
		if strings.ToLower(pilih) == "n" {
			break
		}
	}
	return data
}

func kuitansi(data []string) {
	fmt.Println()
	fmt.Println("Berikut ini adalah daftar pesanan anda :")
	for i, order := range data {
		fmt.Println("Pesanan Ke-", i+1, " : ", order)
	}
	fmt.Println("Terima Kasih atas Pesanan Anda !!")
}

func main() {

	fmt.Println("RAFI COFFE SHOP")
	fmt.Println("========================")

	for _, menu := range menus {
		fmt.Println(menu)
	}
	fmt.Println("======================")
	keranjang()
	kuitansi(data)
	
}