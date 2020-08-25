package main

import (
	"fmt"

	"github.com/livelikeabel/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
