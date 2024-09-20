package main

import "fmt"

func main() {

	tr1 := []int{1, 2, 3}
	tr2 := []int{4, 5, 6}
	tr1 = append(tr1, tr2...) // unpack (~ spread в js)
	fmt.Println(tr1)
	// Упражнение
	var transactions []float64

	for {
		var transaction float64
		fmt.Println("Введите сумму транзакции")
		fmt.Scanf("%f", &transaction) // если не число, то вернет 0
		if transaction == 0 {         // если введено не число (и пока что 0 тоже будет выкидыват из цикла)
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Printf("Transactions: %v", transactions)
}
