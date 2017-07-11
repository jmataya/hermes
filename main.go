package main

import (
	"fmt"

	"github.com/jmataya/hermes/documents"
)

func main() {
	fmt.Println("Hello, world!")

	doc := documents.NewDocument()
	fmt.Printf("%+v\n", doc)
}
