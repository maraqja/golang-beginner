package main

import (
	"fmt"
	"password_app/account" // {название модуля}/{название пакета} - инааче будет искать стандартный пакет
)

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	account1, error := account.NewAccountWithTimeStamp(login, password, url)
	if error != nil {
		fmt.Print(error)
		return
	}
	// account1.generatePassword(10)
	account1.OutputPassword()
	fmt.Println(account1.Url) // букву свойства тоже надо сделать большой в пакете импортируемом, если нужно обращаться к этому свойству тут
	fmt.Println(account1)

}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res) // чтобы при нажатии Enter регало пустую строку
	return res

}
