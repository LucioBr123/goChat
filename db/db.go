package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/LucioBr123/goChat/logger"
)

var (
	db   *sql.DB
	once sync.Once
)

// Configuração do banco
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func getConfig() Config {
	return Config{
		User:     os.Getenv("USUARIO_BANCO"),
		Password: os.Getenv("SENHA_BANCO"),
		Host:     os.Getenv("HOST_BANCO"),
		Port:     os.Getenv("PORTA_BANCO"),
		Database: os.Getenv("BANCO"),
	}
}

// Inicializa a conexao
func initDB() {
	config := getConfig()
	once.Do(func() {
		dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)

		var err error
		db, err = sql.Open("sqlserver", dsn)
		if err != nil {
			logger.SaveLog("Erro ao conectar no banco de dados: " + err.Error())
		}

		//Verifica a conexão
		if err = db.Ping(); err != nil {
			logger.SaveLog("Erro ao verificar a conexão com o banco de dados: " + err.Error())
		}

		logger.SaveLog("Conectado ao banco de dados")
	})
}

func GetDB() *sql.DB {
	initDB()
	return db
}
