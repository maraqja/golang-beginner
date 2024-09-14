package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(`Калькулятор индекса массы тела`)
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.f", imt) // кладем форматированную строку в переменную
	fmt.Print((result))
}

func calculateIMT(userKg float64, userHeight float64) float64 { // объявляем типы входных данных и тип результата
	const IMTPower float64 = 2 // тк функция Pow требует степень в float64

	IMT := userKg / math.Pow(userHeight/100, IMTPower) // неявно преобразуется в float64
	return IMT
}

func getUserInput() (float64, float64) { // Поддерживает возврат нескольких значений из функции, их тип надо описать в ()
	var userKg float64
	var userHeight float64

	fmt.Print("Введите свой рост (в сантиметрах):")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес (в кг):")
	fmt.Scan(&userKg)

	return userKg, userHeight
}
