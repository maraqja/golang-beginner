package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2 // тк функция Pow требует степень в float64

	userHeight := 1.8 // при таком способе объявления переменной нельзя указать тип (всегда будет неявное присвоение) и создать пустую переменную (без присвоения значения, просто создать переменную)
	var userKg float64
	userKg = 100
	IMT := userKg / math.Pow(userHeight, IMTPower) // неявно преобразуется в float64
	fmt.Print(IMT)
}
