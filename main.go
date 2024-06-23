package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/LucioBr123/goChat/logger"
	"github.com/LucioBr123/goChat/routes"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		logger.SaveLog("Error loading .env file")
	}

	r := routes.RegisterRoutes()
	port := os.Getenv("PORTA_SERVIDOR")
	log.Printf("Starting server on :%s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORTA_SERVIDOR")), r))

	log.Printf("Starting server on :%s", port)
}
