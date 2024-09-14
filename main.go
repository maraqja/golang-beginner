package main

import (
	"fmt"
	"math"
)

const IMTPower float64 = 2 // тк функция Pow требует степень в float64

func main() {
	fmt.Println(`Калькулятор индекса массы тела`)
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)

	// МОЖНО ПИСАТЬ ОБЫЧНЫЕ УСЛОВИЯ РАВЕНСТВА СО СВИЧЕМ
	// switch IMT {
	// case 1:
	// 	fmt.Println("У вас сильный недостаток веса")
	// case 2:
	// 	fmt.Println("У вас недостаток веса")
	// }

	// А МОЖНО ПРЯМО УКАЗЫВАТЬ УСЛОВИЯ
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный недостаток веса")
	case IMT < 18.5:
		fmt.Println("У вас недостаток веса")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыток веса")
	default:
		fmt.Println("У вас сильный избыток веса")
	}

	// if IMT < 16 {
	// 	fmt.Println("У вас сильный недостаток веса")
	// } else if IMT < 18.5 {
	// 	fmt.Println("У вас недостаток веса")
	// } else if IMT < 25 {
	// 	fmt.Println("У вас нормальный вес")
	// } else if IMT < 30 {
	// 	fmt.Println("У вас избыток веса")
	// } else {
	// 	fmt.Println("У вас сильный избыток веса")
	// }
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.f", imt) // кладем форматированную строку в переменную
	fmt.Print((result))
}

func calculateIMT(userKg float64, userHeight float64) float64 { // объявляем типы входных данных и тип результата
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
