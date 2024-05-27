package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LucioBr123/goChat/logger"
	"github.com/LucioBr123/goChat/models"
)

type UsuarioRepository struct {
	db *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

// Cria um novo usuário no banco de dados.
func (repo *UsuarioRepository) Create(usuario *models.Usuario) (int64, error) {
	qry := `INSERT 
	          INTO dbo.usuario 
	             ( email
				 , senha
				 , nome
				 , cargo
				 , fotoPath
				 , recado
				 , eAcessoUsuario
				 , eStatusUsuario) 
	        OUTPUT INSERTED.ID
			VALUES (
				   ?
				 , ?
				 , ?
				 , ?
				 , ?
				 , ?
				 , ?
				 , ?)`

	// Executa a query com os valores fornecidos
	var id int64
	err := repo.db.QueryRowContext(context.Background(), qry, usuario.Email, usuario.Senha, usuario.Nome, usuario.Cargo, usuario.FotoPath, usuario.Recado, usuario.EAcessoUsuario, usuario.EStatusUsuario).Scan(&id)
	if err != nil {
		// Loga e retorna o erro
		return 0, logger.SaveLog(fmt.Sprintf("Erro ao criar usuario %s: %v", usuario.Email, err))
	}

	// Loga sucesso
	logger.SaveLog(fmt.Sprintf("Usuário criado com sucesso: %s", usuario.Email))
	return id, nil
}

// Atualiza um usuário no banco de dados.
func (repo *UsuarioRepository) Update(usuario *models.Usuario) error {
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

	// Executa a query com os valores fornecidos
	result, err := repo.db.ExecContext(context.Background(), qry, usuario.Email, usuario.Senha, usuario.Nome, usuario.Cargo, usuario.FotoPath, usuario.Recado, usuario.EAcessoUsuario, usuario.EStatusUsuario, usuario.Id)
	if err != nil {
		// Loga e retorna o erro
		return logger.SaveLog(fmt.Sprintf("Erro ao atualizar usuario %s: %v", usuario.Email, err))
	}

	// Verifica o número de linhas afetadas
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Loga e retorna o erro

		return logger.SaveLog(fmt.Sprintf("Erro ao verificar linhas afetadas ao atualizar usuario %s: %v", usuario.Email, err))
	}
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("Nenhuma linha atualizada para o usuario %s", usuario.Email)
		return logger.SaveLog(errMsg)
	}

	// Loga sucesso
	logger.SaveLog(fmt.Sprintf("Usuário atualizado com sucesso: %s", usuario.Email))
	return nil
}

// Deleta um usuário (soft delete) no banco de dados.
func (repo *UsuarioRepository) Delete(usuarioID int) error {
	qry := `UPDATE dbo.usuario 
	           SET delet = GETDATE()
	         WHERE id = ? 
			   AND delet IS NULL`

	// Executa a query com os valores fornecidos
	result, err := repo.db.ExecContext(context.Background(), qry, usuarioID)
	if err != nil {
		// Loga e retorna o erro
		logger.SaveLog(fmt.Sprintf("Erro ao deletar usuario %d: %v", usuarioID, err))
		return fmt.Errorf("erro ao deletar usuario %d: %w", usuarioID, err)
	}

	// Verifica o número de linhas afetadas
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Loga e retorna o erro
		logger.SaveLog(fmt.Sprintf("Erro ao verificar linhas afetadas ao deletar usuario %d: %v", usuarioID, err))
		return fmt.Errorf("erro ao verificar linhas afetadas ao deletar usuario %d: %w", usuarioID, err)
	}
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("Nenhuma linha deletada para o usuario %d", usuarioID)
		logger.SaveLog(errMsg)
		return fmt.Errorf(errMsg)
	}

	// Loga sucesso
	logger.SaveLog(fmt.Sprintf("Usuário deletado com sucesso: %d", usuarioID))
	return nil
}
