package main

import (
	"fmt"

	"github.com/jmataya/ot-editor/documents"
)

func main() {
	fmt.Println("Hello, world!")

	doc := documents.NewDocument()
	fmt.Printf("%+v\n", doc)
}
