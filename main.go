package main

import (
	"fmt"

	"github.com/livelikeabel/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, error := dictionary.Search("first")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}
}
