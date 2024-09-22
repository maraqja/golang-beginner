package main

import "fmt"

type account struct { // Описываем тип стракта
	login    string
	password string
	url      string
}

func main() {

	// str := "Привет!)"
	str := []rune("Привет!") // Строку можно представить как массив рун
	for _, ch := range str { // цикл по строке, ch имеет тип Rune
		fmt.Println(ch, string(ch))
	}
	// Rune - юникод представления буквы (имеет тип int32)

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1 := account{ // Создаем стракта (инстас типа стракта)
		login:    login,
		password: password,
		url:      url,
	}

	// account2 := account{ // Сокращенная запись, будет зависеть от порядка
	// 	login,
	// 	password,
	// 	url,
	// }

	outputPassword(&account1)

}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scan(&res)
	return res

}

func outputPassword(acc *account) { // struct полностью копируется при передаче в функцию, поэтому есть смысл передавать их по ссылке
	fmt.Println((*acc).login, (*acc).password, acc.url) // GO позволяет использовать shortcode для dereference - вместо (*acc).url использовать просто acc.url
}
