package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LucioBr123/goChat/logger"
	"github.com/LucioBr123/goChat/models"
)

type UsuarioRepositry struct {
	db *sql.DB
}

// Cria usuario
func (repo *UsuarioRepositry) Create(usuario *models.Usuario) error {
	qry := `INSERT 
	          INTO dbo.usuario 
			     ( email
				 , senha
				 , nome
				 , cargo
				 , fotoPath
				 , recado
				 , eAcessoUsuario
				 , eStatusUsuario
				 ) VALUES 
				 ( ?
				 , ?
				 , ?
				 , ?
				 , ?
				 , ?
				 , ?
				 , ?)`

	// Executa a qry com os valores fornecidos
	_, err := repo.db.Exec(qry, usuario.Email, usuario.Senha, usuario.Nome, usuario.Cargo, usuario.FotoPath, usuario.Recado, usuario.EAcessoUsuario, usuario.EStatusUsuario)
	if err != nil {
		//Retorna erro
		return logger.SaveLog(fmt.Sprintf("Erro ao criar usuario %s: %v", usuario.Email, err))
	}

	//Retorna nil
	return nil
}

// Update atualiza o usuario
func (repo *UsuarioRepositry) Update(usuario *models.Usuario) error {
	qry := `UPDATE dbo.usuario 
	           SET email = ?
			     , senha = ?
				 , nome = ?
				 , cargo = ?
				 , fotoPath = ?
				 , recado = ?
				 , eAcessoUsuario = ?
				 , eStatusUsuario = ?
			 WHERE id = ?
			   AND delet IS NULL`

	// Executa a  qry
	result, err := repo.db.ExecContext(context.Background(), qry, usuario.Email, usuario.Senha, usuario.Nome, usuario.Cargo, usuario.FotoPath, usuario.Recado, usuario.EAcessoUsuario, usuario.EStatusUsuario, usuario.Id)
	if err != nil {
		//Retorna erro
		return logger.SaveLog(fmt.Sprintf("Erro ao atualizar usuario %s: %v", usuario.Email, err))
	}

	// Verifica linhas afetadas
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return logger.SaveLog(fmt.Sprintf("Erro ao atualizar usuario, nenhum registro alterado %s: %v", usuario.Email, err))
	}

	//Retorna nil
	return nil
}

func (repo *UsuarioRepositry) Delete(usuario *models.Usuario) error {
	qry := `UPDATE dbo.usuario 
	           SET delet = GETDATE()
			 WHERE id = ?
			   AND delet IS NULL`

	// Executa a  qry
	result, err := repo.db.ExecContext(context.Background(), qry, usuario.Email, usuario.Senha, usuario.Nome, usuario.Cargo, usuario.FotoPath, usuario.Recado, usuario.EAcessoUsuario, usuario.EStatusUsuario, usuario.Id)
	if err != nil {
		//Retorna erro
		return logger.SaveLog(fmt.Sprintf("Erro ao atualizar usuario %s: %v", usuario.Email, err))
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return logger.SaveLog(fmt.Sprintf("Erro ao atualizar usuario, nenhum registro alterado %s: %v", usuario.Email, err))
	}

	//Retorna nil
	return nil
}
