package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func validaDir() error {
	// Verifica se o diretório "log" existe
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		// Cria o diretório "log" se não existir
		if err := os.Mkdir("log", os.ModePerm); err != nil {
			return err
		}
	}

	ano := time.Now().Year()
	anoDir := fmt.Sprintf("log/%d", ano)
	if _, err := os.Stat(anoDir); os.IsNotExist(err) {
		// Cria o diretório do ano se não existir
		if err := os.Mkdir(anoDir, os.ModePerm); err != nil {
			return err
		}
	}

	mes := int(time.Now().Month())
	mesDir := fmt.Sprintf("log/%d/%02d", ano, mes)
	if _, err := os.Stat(mesDir); os.IsNotExist(err) {
		// Cria o diretório do mês se não existir
		if err := os.Mkdir(mesDir, os.ModePerm); err != nil {
			return err
		}
	}

	dia := time.Now().Day()
	diaDir := fmt.Sprintf("log/%d/%02d/%02d.txt", ano, mes, dia)
	if _, err := os.Stat(diaDir); os.IsNotExist(err) {
		// Cria o arquivo do dia se não existir
		if _, err := os.Create(diaDir); err != nil {
			return err
		}
	}

	return nil
}

func saveLog(errorMessage string) {
	if err := validaDir(); err != nil {
		fmt.Println("Erro ao validar diretórios:", err)
		return
	}

	// Obtém a informação de onde a função foi chamada
	_, filename, _, _ := runtime.Caller(1)
	packageName := filepath.Base(filepath.Dir(filename))

	// Obtém parâmetros para o caminho do arquivo de log
	ano := time.Now().Year()
	mes := int(time.Now().Month())
	dia := time.Now().Day()
	caminhoLog := fmt.Sprintf("log/%d/%02d/%02d.txt", ano, mes, dia)

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
