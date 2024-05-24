package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/LucioBr123/goChat/logger"
)

var DB *sql.DB

// Connect cria e retorna uma conexao com o banco de dados sql server
func Connect() error {
	var err error
	// TODO: Arrumar esse user espaço id que tá errado eu acho
	DB, err := sql.Open("sqlserver", fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
		os.Getenv("SERVIDOR_SQL"),
		os.Getenv("USUARIO_BANDO"),
		os.Getenv("SENHA_BANCO"),
		os.Getenv("BANCO")))

	if err != nil {
		return err
	}

	//Configuração de Pool
	// Verifica se a conexão está estabelecida
	if err = DB.PingContext(context.Background()); err != nil {
		DB.Close()
		return err
	}

	maxConAtiva := 25 //Define um maximo de conexão
	if envValue := os.Getenv("MAX_CONEXAO_ABERTA"); envValue != "" {
		maxConAtiva, _ = strconv.Atoi(envValue)
	}

	maxConInativa := 25 //Define um maximo de conexão inativa
	if envValue := os.Getenv("MAX_CONEXAO_INATIVA"); envValue != "" {
		maxConInativa, _ = strconv.Atoi(envValue)
	}

	maxMinInativa := 5 // Define periodo maximo de inatividade
	if envValue := os.Getenv("MAX_CONEXAO_INATIVA"); envValue != "" {
		maxMinInativa, _ = strconv.Atoi(envValue)
	}

	DB.SetMaxOpenConns(maxConAtiva)
	DB.SetMaxIdleConns(maxConInativa)
	DB.SetConnMaxLifetime(time.Duration(maxMinInativa) * time.Minute)

	return nil
}

// Executa uma atualização no banco
func UpdateData(ctx context.Context, db *sql.DB, query string) (bool, error) {
	if ctx == nil {
		return false, logger.SaveLog("context can't be nil")
	}

	if db == nil {
		return false, logger.SaveLog("database can't be nil")
	}

	// Executa a consulta de atualização de dados
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		return false, logger.SaveLog(fmt.Sprintf("error executing query: %v", err))
	}

	// Obtem numero de linhas afetadas
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, logger.SaveLog(fmt.Sprintf("error getting affected rows: %v", err))
	}

	// Retorna True se afetou alguma linha
	return (affectedRows > 0), nil
}

func QryOpen(query string) (*sql.Rows, error) {
	if DB == nil {
		return nil, logger.SaveLog("database is not connected")
	}

	rows, err := DB.QueryContext(context.Background(), query)
	if err != nil {
		return nil, logger.SaveLog(fmt.Sprintf("error executing query: %v", err))
	}
	return rows, nil
}

func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
