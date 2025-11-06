package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("Calculate your profits!")

// 	// var revenue float64
// 	// var expenses float64
// 	// var taxRate float64

// 	revenue, err1 := getUserInfo("Enter your reveneue: ")
// 	expenses, err2 := getUserInfo("Enter your expense: ")
// 	taxRate, err3 := getUserInfo("Enter your tax rate: ")

// 	if err1 != nil {
// 		fmt.Println("Revenue error:", err1)
// 		return
// 	}
// 	if err2 != nil {
// 		fmt.Println("Expenses error:", err2)
// 		return
// 	}
// 	if err3 != nil {
// 		fmt.Println("Tax rate error:", err3)
// 		return
// 	}

// 	calculateProfit(revenue, expenses, taxRate)
// }

// func getUserInfo(infoText string) (float64, error) {
// 	var userInput float64
// 	fmt.Print(infoText)
// 	fmt.Scan(&userInput)

// 	if userInput <= 0 {
// 		return 0, errors.New("value must not be 0 or negative")
// 	}
// 	return userInput, nil
// }

// func calculateProfit(revenue, expenses, taxRate float64) {

// 	fmt.Println("------Result-----")
// 	EBT := revenue - expenses
// 	fmt.Printf("Your earnings without taxes: %v \n", EBT)

// 	profit := EBT * (1 - taxRate/100)
// 	fmt.Printf("Your profits are: %v \n", profit)

// 	ratio := EBT / profit
// 	fmt.Printf("Your total ration: %.2f\n", ratio)

// 	result := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", EBT, profit, ratio)
// 	os.WriteFile("profit_result.txt", []byte(result), 0644)
// }
