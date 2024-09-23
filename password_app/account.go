package main // в одном пакете, просто в разных файлов - это позволяет в другом файле пакета использовать все что описано в этом так, как будто бы оно в одном; это без всяких импортов

// Теперь приложение состоит из двух файлов, поэтому go run main.go будет вызывать ошибку, нужно go run . для запуска всего проекта

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
