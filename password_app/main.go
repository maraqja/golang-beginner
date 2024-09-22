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

func newAccount(login, password, url string) *account { // Это типо можно использовать как конструктор стракта (название просто по соглашению)
	// Хорошо в случае если нужно валидировать входные данные
	return &account{ // возвращаем именно ссылку на структуру, тк иначе мы создадим структуру в функции и еще потом скопириуем в переменную при вызове функции
		login:    login,
		password: password,
		url:      url,
	}
}

// Указываем что функция outputPassword - метод struct account
func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
	acc.url = "test"
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_*!")

func main() {
	login := promptData("Введите логин")
	url := promptData("Введите URL")

	account1 := newAccount(login, "", url)
	account1.generatePassword(10)
	account1.outputPassword()
	fmt.Println(account1.url)

}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scan(&res)
	return res

}
