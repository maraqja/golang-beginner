package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight = 1.8 // неявно присваивается float64 (выведение типов)
	var userKg float64 = 100
	var IMT = userKg / math.Pow(userHeight, 2) // неявно преобразуется в float64
	fmt.Print(IMT)
}
