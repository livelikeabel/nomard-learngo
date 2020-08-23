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

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(123, 4234))

	totalLength, upperName := lenAndUpper("abel")
	fmt.Println(totalLength, upperName)

	repeatMe("abel", "kamake", "jin", "soo", "solomon")
}
