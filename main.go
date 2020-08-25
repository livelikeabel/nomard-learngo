package main

import (
	"fmt"

	"github.com/livelikeabel/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("abel")
	account.Deposit(1000)
	fmt.Println(account)
}
