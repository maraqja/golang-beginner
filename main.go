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
	fmt.Print(`Калькулятор индекса массы тела
Введите свой рост (в сантиметрах):`)
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес (в кг):")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower) // неявно преобразуется в float64
	fmt.Printf("Ваш индекс массы тела: %.f", IMT)
}
