package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	helloStr := "Hello, OTUS!"
	result := stringutil.Reverse(helloStr)
	fmt.Println(result)
}
