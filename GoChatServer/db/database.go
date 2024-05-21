package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/microsoft/go-mssqldb"
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
func UpdateData(query string) (bool, error) {
	// Atualização de dados
	result, err := DB.ExecContext(context.Background(), query)
	if err != nil {
		return false, err
	}

	// Obtem numero de linhas afetadas
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	// Retorna True se afetou alguma linha
	return (affectedRows > 0), nil
}

func QryOpen(query string) (*sql.Rows, error) {
	rows, err := DB.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
