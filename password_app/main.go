package main

import "fmt"

type account struct { // Описываем тип стракта
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1 := account{ // Создаем стракта (инстас типа стракта)
		login:    login,
		password: password,
		url:      url,
	}

	account2 := account{ // Сокращенная запись, будет зависеть от порядка
		login,
		password,
		url,
	}

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res

}

func outputPassword(login, password, url string) {
	fmt.Println(login, password, url)
}
