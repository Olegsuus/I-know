package main

import (
	"encoding/json"
	"fmt"
)

type GoUser struct {
	ID    int
	Name  string
	Email string
}

func main() {
	jsonData := `{"id":1,"name":"Oleg","email":"email@email.ru"}`

	var user GoUser
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Переведено в структуру: ", user)
}
