package main

import (
	"fmt"

	"github.com/LucioBr123/goChat/GoChatServer/tools/logger"
)

func main() {
	// Salva um log de teste
	if err := logger.SaveLog("Este é um erro de teste"); err != nil {
		// Verificar se o erro é diferente de nil
		fmt.Println("Erro ao salvar log:", err)
		if err != nil {
			panic(err)
		}
	}
}
