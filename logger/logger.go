package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func validaDir() error {
	// Obtém parâmetros para o caminho do arquivo de log
	ano := time.Now().Year()
	mes := int(time.Now().Month())
	dia := time.Now().Day()
	diaDir := fmt.Sprintf("log/%d/%02d", ano, mes)

	// Cria diretórios recursivamente
	if err := os.MkdirAll(diaDir, os.ModePerm); err != nil {
		return fmt.Errorf("erro ao criar diretórios: %w", err)
	}

	// Cria o arquivo de log do dia se não existir
	caminhoArquivo := fmt.Sprintf("%s/%02d.txt", diaDir, dia)
	if _, err := os.Stat(caminhoArquivo); os.IsNotExist(err) {
		if _, err := os.Create(caminhoArquivo); err != nil {
			return fmt.Errorf("erro ao criar arquivo de log: %w", err)
		}
	}

	return nil
}

func SaveLog(errorMessage string) {
	if err := validaDir(); err != nil {
		fmt.Println("Erro ao validar diretórios:", err)
		return
	}

	// Obtém a informação de onde a função foi chamada
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Erro ao obter informações do caller")
		return
	}
	packageName := filepath.Base(filepath.Dir(filename))

	// Obtém parâmetros para o caminho do arquivo de log
	ano := time.Now().Year()
	mes := int(time.Now().Month())
	dia := time.Now().Day()
	caminhoLog := fmt.Sprintf("log/%d/%02d/dia-%02d.txt", ano, mes, dia)

	file, err := os.OpenFile(caminhoLog, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
		return
	}
	defer file.Close()

	infoData := time.Now().Format("2006-01-02 15:04:05")

	// Linha a ser adicionada
	novaLinha := fmt.Sprintf("%s - %s - %s\n", infoData, errorMessage, packageName)

	// Adiciona nova linha
	if _, err := fmt.Fprint(file, novaLinha); err != nil {
		fmt.Println("Erro ao escrever no arquivo de log:", err)
		return
	}
}
