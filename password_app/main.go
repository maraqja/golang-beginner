package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct { // Описываем тип стракта
	login    string
	password string
	url      string
}

func newAccount(login, password, urlString string) (*account, error) { // Это типо можно использовать как конструктор стракта (название просто по соглашению)
	// Хорошо в случае если нужно валидировать входные данные
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		// return nil, err
		return nil, errors.New("Invalid URL")
	}
	return &account{ // возвращаем именно ссылку на структуру, тк иначе мы создадим структуру в функции и еще потом скопириуем в переменную при вызове функции
		login:    login,
		password: password,
		url:      urlString,
	}, nil
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

	account1, error := newAccount(login, "", url)
	if error != nil {
		fmt.Print(error)
		return
	}
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
