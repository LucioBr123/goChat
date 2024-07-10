package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/LucioBr123/goChat/logger"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	db   *sql.DB
	once sync.Once
)

// Config holds the database configuration details.
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// getConfig loads the database configuration from environment variables.
func getConfig() Config {
	return Config{
		User:     os.Getenv("USUARIO_BANCO"),
		Password: os.Getenv("SENHA_BANCO"),
		Host:     os.Getenv("HOST_BANCO"),
		Port:     os.Getenv("PORTA_BANCO"),
		Database: os.Getenv("BANCO"),
	}
}

// buildDSN builds the Data Source Name (DSN) string for connecting to the database.
func buildDSN(config Config) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.User, config.Password, config.Host, config.Port, config.Database)
}

// initDB initializes the database connection.
func initDB() error {
	config := getConfig()
	var err error
	once.Do(func() {
		dsn := buildDSN(config)
		db, err = sql.Open("sqlserver", dsn)
		if err != nil {
			logger.LogError(fmt.Sprintf("Erro ao conectar no banco de dados: %v", err))
			return
		}

		// Verifica a conexão
		if err = db.Ping(); err != nil {
			logger.LogError(fmt.Sprintf("Erro ao verificar a conexão com o banco de dados: %v", err))
		}
	})

	return err
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
	_ = initDB()
	return db
}
