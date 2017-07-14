package main

import (
	"fmt"

	"github.com/jmataya/hermes/models"
	"github.com/jmataya/hermes/srv"
)

func main() {
	u := models.User{}
	fmt.Printf("%+v\n", u)
	srv.Run()
}
