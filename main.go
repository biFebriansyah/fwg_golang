package main

import "fmt"

// var password string = "abcd12345"

func main() {
	// variabels()
	// conditional()
	// perluangan()
	// mapsExamp()

	var _, umur = isMature(2005)
	// fmt.Println(check)
	fmt.Println(umur)
}

func isMature(tahunLahir int) (bool, int) {
	if umur := 2023 - tahunLahir; umur > 20 {
		return true, umur
	} else {
		return false, umur
	}
}

func variabels() {

	// manifest type
	var firstName string = "fazz"
	var lastName string = "track"

	var umurs int8 = 127

	fmt.Println(firstName + lastName)
	fmt.Println(umurs)

	// interface
	buah := "mangga"
	fmt.Println(buah)

}

func conditional() {
	username := "user12"
	password := "password"

	if username != "user1" && password != "password" {
		fmt.Println("username salah")

	} else if password != "password" {
		fmt.Println("Password salah")

	} else {
		fmt.Println("Berhasil login")
	}

}

func perluangan() {

	// Array
	var buah = [4]string{
		"apel",
		"mangga",
		"jambu",
		"duren",
	}

	// Slice
	var buahSlice = []string{
		"apel",
		"mangga",
		"jambu",
		"duren",
	}

	buahan1 := buahSlice[0:3]
	fmt.Println(buahan1)

	// forloop
	for i := 0; i < len(buah); i++ {
		fmt.Println(buah[i])
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for index, value := range buah {
		fmt.Println("index = ", index)
		fmt.Println("value = ", value)
	}
}

func mapsExamp() {
	// { key : "value" }
	var _ = map[string]string{
		"1": "satu",
		"2": "dua",
		// "3": 3,
	}

	var talent = []map[string]string{
		{
			"name":  "yunus",
			"class": "fullstack golang",
		},
		{
			"name":  "zaki",
			"class": "fullstack golang",
		},
	}

	for _, value := range talent {
		fmt.Println(value)
	}

	fmt.Println(talent[1]["name"])
}
