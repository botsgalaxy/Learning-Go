package main

import "fmt"

func main() {
	var accountBalance = 1000

	fmt.Println("Welcome to Go Bank")
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
		accountBalance += int(depositAmount)
		fmt.Println("Balance Updated! New amount:", accountBalance)

	} else if choice == 3 {
		var WithdrawalAmount float64
		fmt.Print("Withdrawal amount: ")
		fmt.Scan(&WithdrawalAmount)
		accountBalance -= int(WithdrawalAmount)
		fmt.Println("Balance Updated! New Balance: ", accountBalance)

	} else {
		fmt.Println("Goodbye!")
	}

}
