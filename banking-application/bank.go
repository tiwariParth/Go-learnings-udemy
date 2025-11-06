package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// const accountBalanceFile = "./banking-application/balance.txt"

// func getBalanceToFile() (float64, error) {
// 	data, err := os.ReadFile(accountBalanceFile)
// 	if err != nil {
// 		return 1000, errors.New("failed to find balance")
// 	}
// 	balance, err := strconv.ParseFloat(string(data), 64)

// 	if err != nil {
// 		return 1000, errors.New("failed to parse the stored balance")
// 	}
// 	return balance, nil
// }

// func writeBalanceToFile(balance float64) {
// 	balanceText := fmt.Sprint(balance)
// 	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
// }

// func main() {
// 	var accountBalance, err = getBalanceToFile()

// 	if err != nil {
// 		fmt.Println("ERROR")
// 		fmt.Println(err)
// 		// panic("Can't continue")
// 	}
// 	fmt.Println("Welcom to Go Bank!")
// 	var choice int

// 	for choice != 4 {
// 		fmt.Println("What do you want to do")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit Money")
// 		fmt.Println("3. Withdraw Money")
// 		fmt.Println("4. Exit")
// 		fmt.Print("Your choice: ")
// 		fmt.Scan(&choice)

// 		switch choice {
// 		case 1:
// 			fmt.Println("Your balance is, :", accountBalance)
// 		case 2:
// 			fmt.Print("Your deposit: ")
// 			var depositAmount float64
// 			fmt.Scan(&depositAmount)

// 			if depositAmount <= 0 {
// 				fmt.Println("Invalid number!")
// 				continue
// 			}
// 			accountBalance += depositAmount
// 			writeBalanceToFile(accountBalance)
// 			fmt.Println("Updated balance: ", accountBalance)
// 		case 3:
// 			fmt.Print("Withdrawl amount: ")
// 			var withdrawlAmount float64
// 			fmt.Scan(&withdrawlAmount)
// 			if withdrawlAmount <= 0 {
// 				fmt.Println("Invalid number!")
// 				continue
// 			}
// 			if withdrawlAmount > accountBalance {
// 				fmt.Println("Insufficient Balance! ")
// 				continue
// 			}
// 			accountBalance -= withdrawlAmount
// 			writeBalanceToFile(accountBalance)
// 			fmt.Println("Updated balance: ", accountBalance)
// 		case 4:
// 			// Exit the loop
// 		default:
// 			fmt.Println("Invalid choice! Please try again.")
// 			return
// 		}

// 		// if choice == 1 {
// 		// 	fmt.Println("Your balance is, :", accountBalance)
// 		// } else if choice == 2 {
// 		// 	fmt.Print("Your deposit: ")
// 		// 	var depositAmount float64
// 		// 	fmt.Scan(&depositAmount)

// 		// 	if depositAmount <= 0 {
// 		// 		fmt.Println("Invalid number!")
// 		// 		continue
// 		// 	}

// 		// 	accountBalance += depositAmount
// 		// 	fmt.Println("Updated balance: ", accountBalance)
// 		// } else if choice == 3 {
// 		// 	fmt.Print("Withdrawl amount: ")
// 		// 	var withdrawlAmount float64
// 		// 	fmt.Scan(&withdrawlAmount)

// 		// 	if withdrawlAmount <= 0 {
// 		// 		fmt.Println("Invalid number!")
// 		// 		continue
// 		// 	}
// 		// 	if withdrawlAmount > accountBalance {
// 		// 		fmt.Println("Insufficient Balance! ")
// 		// 		continue
// 		// 	}

// 		// 	accountBalance -= withdrawlAmount
// 		// 	fmt.Println("Updated balance: ", accountBalance)
// 		// }
// 	}
// 	fmt.Println("Thank you for using Go Bank!")

// }
