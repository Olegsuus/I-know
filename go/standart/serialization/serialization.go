package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{
		ID:    1,
		Name:  "Oleg",
		Email: "email@email.ru",
	}

	jsonData, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Данные в виде последовательности байтов:", jsonData)
	fmt.Println("JSON: ", string(jsonData))
}
