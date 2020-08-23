package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func nakedLenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("nakedLenAndUpper func is done!")
	length = len(name)
	uppercase = strings.ToUpper(name)

	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(123, 4234))

	totalLength, upperName := lenAndUpper("abel")
	fmt.Println(totalLength, upperName)

	repeatMe("abel", "kamake", "jin", "soo", "solomon")

	banksaladLength, banksaladUpperName := nakedLenAndUpper("banksalad")
	fmt.Println(banksaladLength, banksaladUpperName)
}
