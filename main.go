package main

import (
	"fmt"
	"go-utils/mystrings"
)

func main() {
	str := mystrings.Reverse("Hello world")
	fmt.Println("[" + str + "]")
}
