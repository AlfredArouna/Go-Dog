package main

import (
	"fmt"

	"github.com/AlfredArouna/Go-Dog/dog"
)

func main() {
	i := 2
	y := dog.Years(i)
	fmt.Println("Humain Age:", i, "Dog Age:", y)
}
