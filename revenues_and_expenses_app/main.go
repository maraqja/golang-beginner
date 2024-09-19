package main

import "fmt"

func main() {

	transactions := [5]int{5, 10, -7, 4, 5} // 3 - размер массива, то что в {} - элементы массива

	transactions_new := transactions // ТУТ ПРОИСХОДИТ ПОЛНОЕ КОПИРОВАНИЕ !!! (НЕ ПО ССЫЛКЕ КАК В JS)
	transactions_new[0] = 200        // ЭТО НЕ ИЗМЕНИТ ИСХОДНЫЙ МАССИВ

	transactions_slice := transactions[0:4] // СЛАЙС ССЫЛАЕТСЯ НА МАССИВ !!!
	transactions_slice[0] = 0               // ЭТО ИЗМЕНИТ ИСХОДНЫЙ МАССИВ

	transactions_slice_new := transactions_slice[0:3] // СЛАЙС ОТ СЛАЙСА ВСЕ РАВНО ССЫЛАЕТСЯ НА ИСХОДНЫЙ МАССИВ
	transactions_slice_new[2] = 100                   // ИЗМЕНИТ В ИСХОДНОМ МАССИВЕ ТАКЖЕ (ну и в слайсе предыдущем тоже)

	fmt.Println(transactions)
	fmt.Printf("new Array: %v\n", transactions_new)
	fmt.Printf("Slice from array: %v\n", transactions_slice)
	fmt.Printf("Slice from slice: %v\n", transactions_slice_new)

	// len (length) - размер слайса
	// cap (capacity) - сколько есть места для элементов в слайсе справа
	fmt.Printf("For slice from array:\nlen: %v\ncap: %v\n", len(transactions_slice), cap(transactions_slice))
	fmt.Printf("For slice from slice:\nlen: %v\ncap: %v\n", len(transactions_slice_new), cap(transactions_slice_new))

	// за счет cap (capacity) можем расширить слайс вправо в рамках размера исхдного массива
	// то есть можно увеличить слайс вправо, но влево - нельзя
	transactions_slice_new = transactions_slice_new[0:5]
	fmt.Printf("Slice from slice: %v\n", transactions_slice_new) // получаем расширенный из самого себя слайс за счет существования capacity
}
