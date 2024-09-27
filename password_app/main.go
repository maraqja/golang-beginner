package main

import (
	"fmt"
	"password_app/account" // {название модуля}/{название пакета} - инааче будет искать стандартный пакет
	"password_app/files"
)

// для добавления внешнего пакета - go get {путь до пакета} - go get github.com/fatih/color
// зависимость от внешних пакетов указана в go.mod
// (зависимости плоские - т.е. если во внешнем импортируемом пакете используется другой пакет, то в go.mod появятся оба пакета)
// восстановить необходимые пакеты по зависимостям - go get

// в go.sum указывает текущий хэш (коммита), относящийся к пакету

// оба файла необходимо коммитить, чтобы можно было восстановить зависимости
func main() {
	files.WriteFile("FILE CONTENT", "kaka.txt")
	files.ReadFile("kaka.txt")
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
