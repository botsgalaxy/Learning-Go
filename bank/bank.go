package main

import "fmt"
import "os"
import "strconv"
import "errors"

func getBalanceFromFile() (float64 , error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil { 
		return 1000,errors.New("failed to find balance file")
	}
	balanceText := string(data)
	balance, err:= strconv.ParseFloat(balanceText, 64)

	if err != nil { 
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)

}

func main() {
	var accountBalance,err = getBalanceFromFile()
	if err != nil { 
		fmt.Println("ERROR:")
		fmt.Println("error occured while reading from file. Using initial balance as 1000")
		fmt.Println("_________________________")
	}

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your Balance is", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance Updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)

		} else if choice == 3 {
			var withdrawalAmount float64
			fmt.Print("Withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount! You can't withdraw more than you have")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance Updated! New Balance: ", accountBalance)
			writeBalanceToFile(accountBalance)

		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank!")

}
