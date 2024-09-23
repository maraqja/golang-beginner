package main

import (
	"fmt"
)

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
