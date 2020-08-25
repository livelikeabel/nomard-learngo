package main

import (
	"fmt"

	"github.com/livelikeabel/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("abel")
	fmt.Println(account)
}
