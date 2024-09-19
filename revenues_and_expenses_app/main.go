package main

import "fmt"

func main() {
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
