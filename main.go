package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2 // тк функция Pow требует степень в float64

	var userHeight float64 // 0.0 если не указать (по умолчанию)
	var userKg float64     // 0.0 если не указать (по умолчанию)
	userKg = 100
	fmt.Println(`Калькулятор индекса массы тела`)
	fmt.Print("Введите свой рост (в сантиметрах):")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес (в кг):")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower) // неявно преобразуется в float64

	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.f", imt) // кладем форматированную строку в переменную
	fmt.Print((result))
}
