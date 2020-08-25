package main

import (
	"fmt"

	"github.com/livelikeabel/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	err := dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search("hello")
	fmt.Println(hello)
}
