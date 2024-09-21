package main

import "fmt"

func main() {
	a := 5
	res := double(a)
	fmt.Println(res)

	pointer_a := &a        // создаем указатель на переменную
	fmt.Println(pointer_a) // выведет адрес переменной в памяти
}

func double(num int) int {
	return num * 2
}
