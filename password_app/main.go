package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a) // Массив != слайс: массив копируется, а слайс передается по ссылке
	fmt.Println(a)
}

// // my
// func reverse(arr *[4]int) { // эта функция будет работать и для слайса, но надо будет переделать немного
// 	left, right := 0, len(arr)-1 // понятно что всегда будет 3

// 	for left < right {
// 		// Меняем местами элементы
// 		arr[left], arr[right] = arr[right], arr[left]
// 		left++
// 		right--
// 	}
// }

func reverse(arr *[4]int) {
	for index, value := range *arr { // ВАЖНО - проходим по значению, а не по ссылкам - если бы по ссылке проходили, то был бы результат как в JS [1,2,2,1]
		(*arr)[len(arr)-1-index] = value // value берем по значению (там примитив)
	}
}
