package main

import (
	"log"

	"github.com/jdferreiro/GoAppSample/bd"
	"github.com/jdferreiro/GoAppSample/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Connectio to DB Fail")
	}
	handlers.Manejadores()
}
