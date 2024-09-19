package main

import "fmt"

func main() {

	transactions := [5]int{5, 10, -7, 4, 5} // 3 - размер массива, то что в {} - элементы массива

	// // Другой способ объявления массива
	// var transactions [3]int
	// transactions[0] = 5
	// transactions[1] = 10
	// transactions[2] = -7

	// // // Объявление пустого массива нужного размера
	// transactions := [3]int{} // 3 - размер массива, то что в {} - элементы массива
	// transactions[0] = 1

	fmt.Println(transactions[1])

	banks := [2]string{"t", "sber"}
	fmt.Println(banks)
	banks[0] = "tinkof"

	// partial_slice := transactions[0:5] // с какого элемента ВКЛЮЧИТЕЛЬНО до какого элемента НЕВКЛЮЧИТЕЛЬНО (или можно понимать как кол-во элементов)
	// partial_slice := transactions[:2] // берем 2 элемента с начала
	partial_slice := transactions[2:] // берем все элементы, после элемента под индексом 2 включительно

	fmt.Println(partial_slice)

}
