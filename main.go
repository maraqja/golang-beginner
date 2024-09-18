package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower float64 = 2 // тк функция Pow требует степень в float64

func main() {
	// // cycles:
	// // 1:
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// // 2: если нужна переменная вне цикла - по сути замена while
	// i := 0
	// for i < 10 {
	// 	fmt.Printf("%d\n", i)
	// 	i++
	// }

	// // 3: бесконечный цикл
	// i := 0
	// for {
	// 	fmt.Printf("%d\n", i)
	// 	i++
	// 	if i == 5 {
	// 		break // Выход из цикла
	// 	}

	// }

	// // 4: range (в блоке по array)

	fmt.Println(`Калькулятор индекса массы тела`)
	// Упражнение
	for {
		userKg, userHeight := getUserInput()
		fmt.Printf("%v, %v\n", userKg, userHeight)
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			fmt.Println(err)
			continue
		}

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

		outputResult(IMT)

		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.f", imt) // кладем форматированную строку в переменную
	fmt.Println((result))
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) { // объявляем типы входных данных и тип результата
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New(("PARAMS_MUST_FLOAT_GREATER_THAN_0"))
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower) // неявно преобразуется в float64
	return IMT, nil                                    // nil - специальное значение, обозначающее отсутствие чего-либо
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

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Println("Хотите продолжить? y/n")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
