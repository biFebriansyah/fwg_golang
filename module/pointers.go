package module

import "fmt"

func PointerExamp() {
	var idUser int = 20
	result := getData(&idUser) // pass by value
	fmt.Println(result)

}

func getData(id *int) string {
	return fmt.Sprintf("User with id %d", id)
}
