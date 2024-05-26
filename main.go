package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LucioBr123/goChat/routes"
)

func main() {
	r := routes.RegisterRoutes()
	log.Printf("Starting server on :%s", os.Getenv("PORTA_SERVIDOR"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORTA_SERVIDOR")), r))
}
