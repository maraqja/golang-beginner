package main

import (
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "Город") // хотим забирать из строки флаги (не работает до вызова parse) - go run . --name = "Kaka"
	format := flag.Int("format", 1, "Формат вывода погоды")
	// isAdmin := flag.Bool("isAdmin", false, "Является ли администратором?")
	flag.Parse() // после описания всех флагов вызываем

	fmt.Println(*city)
	fmt.Println(*format)

}