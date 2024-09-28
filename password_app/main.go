package main

import (
	"fmt"
	"password_app/account" // {название модуля}/{название пакета} - инааче будет искать стандартный пакет
	"password_app/files"
	"password_app/output"
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
		userChoice := promptData([]string{"Создать аккаунт", "Найти аккаунт", "Удалить аккаунт", "Выход", "Выберите вариант"})

		switch userChoice {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL"})
	foundAccounts := vault.FindAccountsByUrl(url)
	// fmt.Println(foundAccounts)
	if len(foundAccounts) == 0 {
		output.PrintError("Не найдено аккаунтов по url = " + url)
	}
	for _, account := range foundAccounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		output.PrintError("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Print(err)
		return
	}

	vault.AddAccount(myAccount)
}

func promptData[T any](prompt []T) string {
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
