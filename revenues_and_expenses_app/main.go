package main

import "fmt"

func main() {

	transactions := [3]int{5, 10, -7} // 3 - размер массива, то что в {} - элементы массива

	// // Другой способ объявления массива
	// var transactions [3]int
	// transactions[0] = 5
	// transactions[1] = 10
	// transactions[2] = -7

	// // // Объявление пустого массива нужного размера
	// transactions := [3]int{} // 3 - размер массива, то что в {} - элементы массива
	// transactions[0] = 1

	banks := [2]string{"t", "sber"}

	banks[0] = "tinkof"
	fmt.Println(transactions[1])
	fmt.Println(banks)
}
