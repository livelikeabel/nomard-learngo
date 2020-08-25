package main

import (
	"fmt"

	"github.com/livelikeabel/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("abel")
	account.Deposit(1000)
	fmt.Println(account.Balance())
	err := account.Withdraw(2000)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.Balance())
}
