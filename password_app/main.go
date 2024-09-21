package main

import "fmt"

func main() {
	a := 5
	double(&a) // это будет МУТИРОВАНИЕ исходного a !!!
	fmt.Println(a)
}

func double(num *int) { // *type -  указатель на переменную типа type
	// должны получить значение переменной по указателю
	*num = *num * 2 // *num - дереференс (dereference) - получаем значение переменной по адресу указателя
}
