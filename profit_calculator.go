package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Calculate your profits!")

// 	// var revenue float64
// 	// var expenses float64
// 	// var taxRate float64

// 	revenue := getUserInfo("Enter your reveneue: ")
// 	expenses := getUserInfo("Enter your expense: ")
// 	taxRate := getUserInfo("Enter your tax rate: ")

// 	calculateProfit(revenue, expenses, taxRate)
// }

// func getUserInfo(infoText string) float64 {
// 	var userInput float64
// 	fmt.Print(infoText)
// 	fmt.Scan(&userInput)
// 	return userInput
// }

// func calculateProfit(revenue, expenses, taxRate float64) {
// 	EBT := revenue - expenses
// 	fmt.Printf("Your earnings without taxes: %v \n", EBT)

// 	profit := EBT * (1 - expenses/100)
// 	fmt.Printf("Your profits are: %v \n", profit)

// 	ratio := EBT / profit
// 	fmt.Printf("Your total ration: %v\n", ratio)
// }
