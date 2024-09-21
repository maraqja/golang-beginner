package main

import "fmt"

func main() {
	//map = key - value
	m := map[string]string{ // [тип ключа, тип значения]
		"PurpleSchool": "https://purpleschool.ru",
	}
	fmt.Println(m)
	fmt.Println(m["PurpleSchool"]) // Читаем элемент мапа
	m["PurpleSchool"] = "https://new.purpleschool.ru"
	fmt.Println(m["PurpleSchool"])
	m["Google"] = "https://google.com" // Добавляем элемент мапа
	m["Yandex"] = "https://yandex.ru"
	fmt.Println(m)
	delete(m, "Yandex") // Удаляем элемент мапа
	fmt.Println(m)
}
