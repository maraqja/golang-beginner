package main

import "fmt"

func main() {

	// tr := []string{}
	// tr = append(tr, "1") // на каждый append будет проходить дополнительное выделение памяти для go
	// tr = append(tr, "2")
	// fmt.Println(tr)

	tr := make([]string, 0, 2) // make позволяет сказать, что создаем массив определенной len (текущая длина) и cap (capacity - сколько мы максимально можем вместить в массив)
	// make (описание типа массива, длина, capacity)
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "1") // теперь не будет происходит выделение памяти на каждый append
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "2") // теперь не будет происходит выделение памяти на каждый append
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "3")          // а на этот уже будет
	fmt.Println(len(tr), cap(tr)) // тут capacity измениля на 4 (а не на 3 как ожидалось)
	tr = append(tr, "4")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "5")
	fmt.Println(len(tr), cap(tr)) // тут capacity снова увеличилось в 2 раза
	// вообщем капасити будет расти в 2 раза при новой len > текущей cap, пока cap < 256 (дальше будет уменьшаться увеличение вплоть до 1.25x)
	// так просто реализовано для оптимизации в самом go - это легко видеть в исходниках го (просто на гите) - метод nextslicecap

	fmt.Println(tr)
	// ------------------------ //
	var transactions []float64 // уточнение по динамическим массивам: под капотом создается массив, на который ссылается наш слайс transactions

	for {
		var transaction float64
		fmt.Println("Введите сумму транзакции")
		fmt.Scanf("%f", &transaction) // если не число, то вернет 0
		if transaction == 0 {         // если введено не число (и пока что 0 тоже будет выкидыват из цикла)
			break
		}
		transactions = append(transactions, transaction)
	}
	// fmt.Printf("Transactions: %v", transactions)

	// Упражнение
	balance := 0.0
	for _, value := range transactions {
		// fmt.Println(index, value)
		balance += value
	}
	fmt.Printf("Balance, %f\n", balance)
}
