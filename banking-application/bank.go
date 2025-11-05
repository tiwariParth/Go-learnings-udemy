package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcom to Go Bank!")
	// fmt.Println("What do you want to do")
	// fmt.Println("1. Check balance")
	// fmt.Println("2. Deposit Money")
	// fmt.Println("3. Withdraw Money")
	// fmt.Println("4. Exit")

	var choice int
	// fmt.Print("Your choice: ")
	// fmt.Scan(&choice)

	for choice != 4 {
		fmt.Println("What do you want to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		// var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is, :", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid number!")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Updated balance: ", accountBalance)
		case 3:
			fmt.Print("Withdrawl amount: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)
			if withdrawlAmount <= 0 {
				fmt.Println("Invalid number!")
				continue
			}
			if withdrawlAmount > accountBalance {
				fmt.Println("Insufficient Balance! ")
				continue
			}
			accountBalance -= withdrawlAmount
			fmt.Println("Updated balance: ", accountBalance)
		case 4:
			// Exit the loop
		default:
			fmt.Println("Invalid choice! Please try again.")
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is, :", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid number!")
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Updated balance: ", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdrawl amount: ")
		// 	var withdrawlAmount float64
		// 	fmt.Scan(&withdrawlAmount)

		// 	if withdrawlAmount <= 0 {
		// 		fmt.Println("Invalid number!")
		// 		continue
		// 	}
		// 	if withdrawlAmount > accountBalance {
		// 		fmt.Println("Insufficient Balance! ")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawlAmount
		// 	fmt.Println("Updated balance: ", accountBalance)
		// }
	}
	fmt.Println("Thank you for using Go Bank!")

}
