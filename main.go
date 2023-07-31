package main

import (
	"biFebriansyah/gointro/module"
	"fmt"
)

// var password string = "abcd12345"

func main() {
	var zaki = module.CreateUser("zaki", "jkt", "jakarta")
	checkAlamat(zaki)

	// var tanggal int = 20
	// var pTanggal int = tanggal

	// fmt.Println("tanggal ", tanggal)
	// fmt.Println("ptanggal ", pTanggal)

	// pTanggal = 40

	// // value ?
	// fmt.Println("tanggal ", tanggal)
	// fmt.Println("ptanggal ", pTanggal)

	cetak(20)

	var respone = map[string]interface{}{
		"name": "julian",
		"age":  20,
	}

	fmt.Println(respone)
}

func cetak(data interface{}) {
	var result = 10 + data.(int)
	fmt.Println(result)
}

func checkAlamat(user module.IfUser) {
	fmt.Println(user.GetName())
}
