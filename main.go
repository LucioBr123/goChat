package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/LucioBr123/goChat/logger"
	"github.com/LucioBr123/goChat/routes"
)

func main() {
	r := routes.RegisterRoutes()
	portStr := os.Getenv("PORTA_SERVIDOR")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.SaveLog(fmt.Sprintf("Error converting port to integer: %v", err))
	}
	log.Printf("Starting server on :%d", port)
	// TODO: Verificar motivo de não conseguir obter a variavel env

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORTA_SERVIDOR")), r))
}
