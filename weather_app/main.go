package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	city := flag.String("city", "", "Город") // хотим забирать из строки флаги (не работает до вызова parse) - go run . --name = "Kaka"
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse() // после описания всех флагов вызываем

	fmt.Println(*city)
	fmt.Println(*format)

	r := strings.NewReader("Привет! Я поток данных") // создаем ридер
	block := make([]byte, 4) // хотим читать по 4 байта из потока - создаем массив байт размера 4, в который будем складывать блоки (определенного размера (4)) из Ридера
	for { // пока не конец потока
		_, err := r.Read(block) // читаем определенное число байт в переменную block
		if err == io.EOF { // ошибка может быть 1) просто ошибка чтения 2) закончился поток
			break;
		}
		fmt.Printf("%q\n", block)
	}

} 