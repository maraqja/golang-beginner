package main

import "fmt"

type stringMap = map[string]string

func main() {
	m := make(stringMap, 2)
	m["A"] = "1"
	m["B"] = "2"
	m["C"] = "3" // на этот уже будет выделяться память при добавлении
	fmt.Println(m)
	fmt.Println(len(m)) // у мапа нет capacity

	// var bookmarks map[string]string
	bookmarks := stringMap{}
	fmt.Println("Закладки")
menu:
	for {
		user_choice := getUserChoice()

		switch user_choice {
		case 1:
			getBookmarks(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
		case 3:
			bookmarks = deleteBookmark(bookmarks)
		case 4:
			break menu // выходим именно из цикла, а не из case
		}
	}
}

func getUserChoice() int {
	var user_choice int
	fmt.Println("Выберите")
	fmt.Println("1 - Посмотреть закладки")
	fmt.Println("2 - Добавить закладку")
	fmt.Println("3 - Удалить закладку")
	fmt.Println("4 - Выход")
	fmt.Scan(&user_choice)
	return user_choice
}

func getBookmarks(bookmarks stringMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Закладок нет")
	}
	for key, value := range bookmarks {
		fmt.Printf("%s : %s\n", key, value)
	}
}

func addBookmark(bookmarks stringMap) stringMap {
	var new_bookmark_key, new_bookmark_value string
	fmt.Println("Введите название: ")
	fmt.Scan(&new_bookmark_key)
	fmt.Println("Введите ссылку: ")
	fmt.Scan(&new_bookmark_value)
	bookmarks[new_bookmark_key] = new_bookmark_value
	return bookmarks
}

func deleteBookmark(bookmarks stringMap) stringMap {
	var bookmark_key string
	fmt.Println("Введите название: ")
	fmt.Scan(&bookmark_key)
	delete(bookmarks, bookmark_key)
	return bookmarks
}
