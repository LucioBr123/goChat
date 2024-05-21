package connection

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

func Connect() (*sql.DB, error) {
	// Conecta ao banco de dados
	db, err := sql.Open("sqlserver", "server="+os.Getenv("SERVIDOR_SQL")+";user id="+os.Getenv("USUARIO_BANCO")+";password="+os.Getenv("SENHA_BANCO")+";database="+os.Getenv("BANCO"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func UpdateData(db *sql.DB, query string) (bool, error) {
	// Atualização de dados
	result, err := db.ExecContext(context.Background(), query)
	if err != nil {
		return false, err
	}

	// Obtem numero de linhas afetadas
	affectedRows, _ := result.RowsAffected()
	if err != nil {
		return false, err
	}

	// Retorna True se afetou alguma linha
	return (affectedRows > 0), nil
}
