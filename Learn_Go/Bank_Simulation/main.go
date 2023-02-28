package main

import (
	"fmt"

	"github.com/mingeun3669/Bank_Simulation/accounts"
)

func main() {
	var owner string
	fmt.Print("Owner's Name ? ")
	fmt.Scan(&owner)
	account := accounts.NewAccount(owner)
	fmt.Println(account)

	account.Deposit(999)
	fmt.Println(account.Balance())

	err := account.Withdraw(9999)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
