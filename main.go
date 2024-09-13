package main

func main() {
	var userHeight = 1.8          // неявно присваивается float64 (выведение типов)
	var userKg = 100              // неявно присваивается int (выведение типов)
	var IMT = userKg / userHeight // нельзя разделить - в go нет неявного преобразования
}
