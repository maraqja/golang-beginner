package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct { // Описываем тип стракта
	login    string
	password string
	url      string
}

func newAccount(login, password, urlString string) (*account, error) { // Это типо можно использовать как конструктор стракта (название просто по соглашению)
	// Хорошо в случае если нужно валидировать входные данные
	if login == "" { // логин не был передан: дефолтное значение непереданноо строки - пустая строка
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		// return nil, err
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &account{ // возвращаем именно ссылку на структуру, тк иначе мы создадим структуру в функции и еще потом скопириуем в переменную при вызове функции
		login:    login,
		password: password,
		url:      urlString,
	}
	if password == "" { // Если не передали пароль, то генерим
		newAcc.generatePassword(10)
	}
	return newAcc, nil
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

// В GO нет как такового наследования, но его заменяет композиция
type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account   // Встраивание - появляются embeded fields
}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) { // Это типо можно использовать как конструктор стракта (название просто по соглашению)
	// Хорошо в случае если нужно валидировать входные данные
	if login == "" { // логин не был передан: дефолтное значение непереданноо строки - пустая строка
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		// return nil, err
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimeStamp{ // возвращаем именно ссылку на структуру, тк иначе мы создадим структуру в функции и еще потом скопириуем в переменную при вызове функции
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{ // Приходится объявлять так с помощью встраивания
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" { // Если не передали пароль, то генерим
		// newAcc.account.generatePassword(10) // ориг запись
		newAcc.generatePassword(10)
	}
	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_*!")

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	account1, error := newAccountWithTimeStamp(login, password, url)
	if error != nil {
		fmt.Print(error)
		return
	}
	// account1.generatePassword(10)
	account1.outputPassword()
	fmt.Println(account1.url)
	fmt.Println(account1)

}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res) // чтобы при нажатии Enter регало пустую строку
	return res

}
