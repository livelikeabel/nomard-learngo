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

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrinkWithSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(multiply(123, 4234))

	totalLength, upperName := lenAndUpper("abel")
	fmt.Println(totalLength, upperName)

	repeatMe("abel", "kamake", "jin", "soo", "solomon")

	banksaladLength, banksaladUpperName := nakedLenAndUpper("banksalad")
	fmt.Println(banksaladLength, banksaladUpperName)

	totalResult := superAdd(10, 20, 30, 40, 50)
	fmt.Println(totalResult)

	fmt.Println(canIDrink(16))
}
