package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct { // Описываем тип стракта
	login    string
	password string
	url      string
}

// Указываем что функция outputPassword - метод struct account
func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
	acc.url = "test"
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_*!")

func main() {
	login := promptData("Введите логин")
	url := promptData("Введите URL")

	account1 := account{ // Создаем стракта (инстас типа стракта)
		login:    login,
		password: generatePassword(10),
		url:      url,
	}

	// account2 := account{ // Сокращенная запись, будет зависеть от порядка
	// 	login,
	// 	password,
	// 	url,
	// }
	account1.outputPassword()
	fmt.Println(account1.url)

}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scan(&res)
	return res

}

func generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(res)
}
