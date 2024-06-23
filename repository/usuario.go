package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LucioBr123/goChat/logger"
	"github.com/LucioBr123/goChat/models"
)

// UsuarioRepository represents the repository for user operations.
type UsuarioRepository struct {
	db *sql.DB
}

// NewUsuarioRepository creates a new instance of UsuarioRepository.
func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

// Create inserts a new user into the database.
func (repo *UsuarioRepository) Create(ctx context.Context, usuario *models.Usuario) (int64, error) {
	qry := `INSERT INTO dbo.usuario 
	         (email, senha, nome, cargo, fotoPath, recado, eAcessoUsuario, eStatusUsuario) 
	         OUTPUT INSERTED.ID
	         VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	var id int64
	err := repo.db.QueryRowContext(ctx, qry, usuario.Email, usuario.Senha, usuario.Nome, usuario.Cargo, usuario.FotoPath, usuario.Recado, usuario.EAcessoUsuario, usuario.EStatusUsuario).Scan(&id)
	if err != nil {
		return 0, logger.SaveLog(fmt.Sprintf("Erro ao criar usuario %s: %v", usuario.Email, err))
	}

	logger.SaveLog(fmt.Sprintf("Usuário criado com sucesso: %s", usuario.Email))
	return id, nil
}

// Update modifies an existing user in the database.
func (repo *UsuarioRepository) Update(ctx context.Context, usuario *models.Usuario) error {
	qry := `UPDATE dbo.usuario 
	        SET email = ?, senha = ?, nome = ?, cargo = ?, fotoPath = ?, recado = ?, eAcessoUsuario = ?, eStatusUsuario = ? 
	        WHERE id = ? AND delet IS NULL`

	result, err := repo.db.ExecContext(ctx, qry, usuario.Email, usuario.Senha, usuario.Nome, usuario.Cargo, usuario.FotoPath, usuario.Recado, usuario.EAcessoUsuario, usuario.EStatusUsuario, usuario.Id)
	if err != nil {
		return logger.SaveLog(fmt.Sprintf("Erro ao atualizar usuario %s: %v", usuario.Email, err))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return logger.SaveLog(fmt.Sprintf("Erro ao verificar linhas afetadas ao atualizar usuario %s: %v", usuario.Email, err))
	}
	if rowsAffected == 0 {
		return logger.SaveLog(fmt.Sprintf("Nenhuma linha atualizada para o usuario %s", usuario.Email))
	}

	logger.SaveLog(fmt.Sprintf("Usuário atualizado com sucesso: %s", usuario.Email))
	return nil
}

// Delete performs a soft delete on a user in the database.
func (repo *UsuarioRepository) Delete(ctx context.Context, usuarioID int) error {
	qry := `UPDATE dbo.usuario 
	        SET delet = GETDATE() 
	        WHERE id = ? AND delet IS NULL`

	result, err := repo.db.ExecContext(ctx, qry, usuarioID)
	if err != nil {
		logger.SaveLog(fmt.Sprintf("Erro ao deletar usuario %d: %v", usuarioID, err))
		return fmt.Errorf("erro ao deletar usuario %d: %w", usuarioID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.SaveLog(fmt.Sprintf("Erro ao verificar linhas afetadas ao deletar usuario %d: %v", usuarioID, err))
		return fmt.Errorf("erro ao verificar linhas afetadas ao deletar usuario %d: %w", usuarioID, err)
	}
	if rowsAffected == 0 {
		return logger.SaveLog(fmt.Sprintf("Nenhuma linha deletada para o usuario %d", usuarioID))
	}

	logger.SaveLog(fmt.Sprintf("Usuário deletado com sucesso: %d", usuarioID))
	return nil
}
