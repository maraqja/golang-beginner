package main

import (
	"fmt"
	"password_app/account" // {название модуля}/{название пакета} - инааче будет искать стандартный пакет
	"password_app/encrypter"
	"password_app/files"
	"password_app/output"
	"strings"

	"github.com/joho/godotenv"
)

// для добавления внешнего пакета - go get {путь до пакета} - go get github.com/fatih/color
// зависимость от внешних пакетов указана в go.mod
// (зависимости плоские - т.е. если во внешнем импортируемом пакете используется другой пакет, то в go.mod появятся оба пакета)
// восстановить необходимые пакеты по зависимостям - go get

// в go.sum указывает текущий хэш (коммита), относящийся к пакету

// оба файла необходимо коммитить, чтобы можно было восстановить зависимости

var menu = map[string]func(*account.VaultWithDb){ // в каждом ключе map - функция
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuChoices = []string{"Создать аккаунт", "Найти аккаунт по URL", "Найти аккаунт по логину", "Удалить аккаунт", "Выход", "Выберите вариант"}

func main() {
	err := godotenv.Load() // забираем переменные из файлы и засовываем их в окружение
	if err != nil {
		output.PrintError("Не удалось найти env файл")
	}
	db := files.NewJsonDb("data.json")
	// db := cloud.NewCloudDb("https://kaka.com") // теперь можем использовать любую struct, реализующую интерфейса бд
	enc := encrypter.NewEncrypter()
	vault := account.NewVault(db, enc)
Menu:
	for {
		userChoice := promptData(menuChoices...) // передаем в функцию элементы слайса через запятую
		menuFunc := menu[userChoice]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}

}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL")
	foundAccounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	if len(foundAccounts) == 0 {
		output.PrintError("Не найдено аккаунтов по url = " + url)
	}
	for _, account := range foundAccounts {
		account.Output()
	}
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите login")
	foundAccounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	if len(foundAccounts) == 0 {
		output.PrintError("Не найдено аккаунтов по login = " + login)
	}
	for _, account := range foundAccounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		output.PrintError("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Print(err)
		return
	}

	vault.AddAccount(myAccount)
}

func promptData(prompt ...string) string { // сложит все аргументы в слайс prompt
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Printf("%v. %v\n", i+1, line)
		}
	}
	var res string
	fmt.Scanln(&res) // чтобы при нажатии Enter регало пустую строку
	return res

}
