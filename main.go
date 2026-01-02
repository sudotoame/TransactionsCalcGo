package main

import (
	"fmt"
)

func main() {
	transactions := []float64{}
	for {
		newTransiction, err := scanTransations()
		if newTransiction == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		transactions = append(transactions, newTransiction)
		fmt.Println("Транзакции: ", transactions)
	}
	balance := calculateBalance(transactions)
	fmt.Printf("Ваш баланса: %.2f\n", balance)
}

func calculateBalance(transactions []float64) (balance float64) {
	for _, value := range transactions {
		balance += value
	}
	return balance
}

func scanTransations() (digit float64, err error) {
	fmt.Print("Введите транзакцию (n для выхода): ")
	_, err = fmt.Scan(&digit)
	if err != nil {
		return digit, err
	}
	return digit, nil
}
