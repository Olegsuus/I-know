package main

// Сериализация - это процесс преобразования объекта или структуры данных в последовательность байтов или текстовый формат
// (например JSON или XML) для сохранения в файл, передачи по сети или хранения в базе данных

//Пример сериализации данных

//type User struct {
//	ID    int    `json:"id"`
//	Name  string `json:"name"`
//	Email string `json:"email"`
//}
//
//func main() {
//	user := User{
//		ID:    1,
//		Name:  "Oleg",
//		Email: "email@email.ru",
//	}
//
//	jsonData, err := json.Marshal(&user)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Данные в виде последовательности байтов:", jsonData)
//	fmt.Println("JSON: ", string(jsonData))
//}
