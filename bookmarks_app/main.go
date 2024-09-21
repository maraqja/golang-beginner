package main

import "fmt"

func main() {
	// var bookmarks map[string]string
	bookmarks := map[string]string{}
	fmt.Println("Закладки")
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
			break
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

func getBookmarks(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("Закладок нет")
	}
	for key, value := range bookmarks {
		fmt.Printf("%s : %s\n", key, value)
	}
}

func addBookmark(bookmarks map[string]string) map[string]string {
	var new_bookmark_key, new_bookmark_value string
	fmt.Println("Введите название: ")
	fmt.Scan(&new_bookmark_key)
	fmt.Println("Введите ссылку: ")
	fmt.Scan(&new_bookmark_value)
	bookmarks[new_bookmark_key] = new_bookmark_value
	return bookmarks
}

func deleteBookmark(bookmarks map[string]string) map[string]string {
	var bookmark_key string
	fmt.Println("Введите название: ")
	fmt.Scan(&bookmark_key)
	delete(bookmarks, bookmark_key)
	return bookmarks
}
