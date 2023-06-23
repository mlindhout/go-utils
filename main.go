package main

import (
	"fmt"
	"github.com/mlindhout/go-utils/mystrings"
)

func main() {
	str := mystrings.Reverse("Hello world")
	fmt.Println("[" + str + "]")
}
