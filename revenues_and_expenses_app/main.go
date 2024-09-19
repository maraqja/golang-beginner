package main

import "fmt"

func main() {
	// Динамический массив - массив, которому не задан размер
	transactions := []int{5, 10, -7, 4, 5} // создаем массив без ограничения по длине (по факту это слайс под капотом, хранящий ссылку на массив (хз какой))
	slice := transactions                  // тут по сути копируем слайс => изменение в slice изменит исходный слайс (который типо динамический массив)
	slice[0] = 0                           // изменится в transactions

	fmt.Printf("Dynamic array (it is slice !!!): %v\n", transactions)
	fmt.Printf("Slice from dynamic array: %v\n", slice)
	// slice[6] = 200                         // будет вызывать ошибку, тк нет такого индекса в массиве
	// нужно использовать append для вставки новых элементов
	transactions = append(transactions, 100) // append можно использовать только на slice-ах
	// ВАЖНО - ссылающийся slice не изменится: APPEND создает новый массив в памяти, slice будет ссылаться на старые данные transactions
	fmt.Printf("Dynamic array (it is slice !!!): %v\n", transactions)
	fmt.Printf("Slice from dynamic array: %v\n", slice)
}
