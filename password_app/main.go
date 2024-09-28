package main

import (
	"fmt"
	"password_app/account" // {название модуля}/{название пакета} - инааче будет искать стандартный пакет
	"password_app/files"

	"github.com/fatih/color"
)

// для добавления внешнего пакета - go get {путь до пакета} - go get github.com/fatih/color
// зависимость от внешних пакетов указана в go.mod
// (зависимости плоские - т.е. если во внешнем импортируемом пакете используется другой пакет, то в go.mod появятся оба пакета)
// восстановить необходимые пакеты по зависимостям - go get

// в go.sum указывает текущий хэш (коммита), относящийся к пакету

// оба файла необходимо коммитить, чтобы можно было восстановить зависимости
func main() {
	db := files.NewJsonDb("data.json")
	// db := cloud.NewCloudDb("https://kaka.com") // теперь можем использовать любую struct, реализующую интерфейса бд
	vault := account.NewVault(db)
Menu:
	for {
		userChoice := getMenu()
		switch userChoice {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func getMenu() int {
	var userChoice int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4 (и любое другое). Выход")
	fmt.Scan(&userChoice)
	return userChoice
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL")
	foundAccounts := vault.FindAccountsByUrl(url)
	// fmt.Println(foundAccounts)
	if len(foundAccounts) == 0 {
		color.Red("Не найдено аккаунтов по url = %s", url)
	}
	for _, account := range foundAccounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
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

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res) // чтобы при нажатии Enter регало пустую строку
	return res

}
