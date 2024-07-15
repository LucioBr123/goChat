package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LucioBr123/goChat/logger"
	"github.com/LucioBr123/goChat/models"
)

// UsuarioRepository representa o repositório para operações de usuário.
type UsuarioRepository struct {
	db *sql.DB
}

// NewUsuarioRepository cria uma nova instância de UsuarioRepository.
func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

// Create insere um novo usuário no banco de dados.
func (repo *UsuarioRepository) Create(ctx context.Context, usuario *models.Usuario) (int64, error) {
	qry := `
		INSERT INTO dbo.usuario (
			email, senha, nome, cargo, fotoPath, recado, eAcessoUsuario, eStatusUsuario
		) 
		OUTPUT INSERTED.ID
		VALUES (
			@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8
		)`

	var id int64
	err := repo.db.QueryRowContext(ctx, qry,
		sql.Named("p1", usuario.Email),
		sql.Named("p2", usuario.Senha),
		sql.Named("p3", usuario.Nome),
		sql.Named("p4", usuario.Cargo),
		sql.Named("p5", usuario.FotoPath),
		sql.Named("p6", usuario.Recado),
		sql.Named("p7", usuario.EAcessoUsuario),
		sql.Named("p8", usuario.EStatusUsuario),
	).Scan(&id)
	if err != nil {
		logger.LogError(fmt.Sprintf("Erro ao criar usuario %s: %v", usuario.Email, err))
		return 0, err
	}

	logger.LogError(fmt.Sprintf("Usuário criado com sucesso: %s", usuario.Email))
	return id, nil
}

// Update modifica um usuário existente no banco de dados.
func (repo *UsuarioRepository) Update(ctx context.Context, usuario *models.Usuario) error {
	qry := `UPDATE dbo.usuario
	           SET email = @p1
			     , senha = @p2
				 , nome  = @p3
				 , cargo = @p4
				 , fotoPath = @p5
				 , recado = @p6
				 , eAcessoUsuario = @p7
				 , eStatusUsuario = @p8
		    WHERE id = @p9 
			  AND delet IS NULL`

	result, err := repo.db.ExecContext(ctx, qry,
		sql.Named("p1", usuario.Email),
		sql.Named("p2", usuario.Senha),
		sql.Named("p3", usuario.Nome),
		sql.Named("p4", usuario.Cargo),
		sql.Named("p5", usuario.FotoPath),
		sql.Named("p6", usuario.Recado),
		sql.Named("p7", usuario.EAcessoUsuario),
		sql.Named("p8", usuario.EStatusUsuario),
		sql.Named("p9", usuario.Id),
	)
	if err != nil {
		return logger.LogError(fmt.Sprintf("Erro ao atualizar usuario %s: %v", usuario.Email, err))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return logger.LogError(fmt.Sprintf("Erro ao verificar linhas afetadas ao atualizar usuario %s: %v", usuario.Email, err))
	}
	if rowsAffected == 0 {
		return logger.LogError(fmt.Sprintf("Nenhuma linha atualizada para o usuario %s", usuario.Email))
	}

	logger.LogError(fmt.Sprintf("Usuário atualizado com sucesso: %s", usuario.Email))
	return nil
}

// Delete realiza um soft delete em um usuário no banco de dados.
func (repo *UsuarioRepository) Desativa(ctx context.Context, usuarioID int) error {
	qry := `UPDATE dbo.usuario 
	          SET delet = GETDATE() 
	        WHERE id = @p1 
			  AND delet IS NULL`

	result, err := repo.db.ExecContext(ctx, qry, sql.Named("p1", usuarioID))
	if err != nil {
		logger.LogError(fmt.Sprintf("Erro ao deletar usuario %d: %v", usuarioID, err))
		return fmt.Errorf("erro ao deletar usuario %d: %w", usuarioID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.LogError(fmt.Sprintf("Erro ao verificar linhas afetadas ao deletar usuario %d: %v", usuarioID, err))
		return fmt.Errorf("erro ao verificar linhas afetadas ao deletar usuario %d: %w", usuarioID, err)
	}
	if rowsAffected == 0 {
		return logger.LogError(fmt.Sprintf("Nenhuma linha deletada para o usuario %d", usuarioID))
	}

	logger.LogError(fmt.Sprintf("Usuário deletado com sucesso: %d", usuarioID))
	return nil
}

// Ativa ativa um usuário no banco de dados.
func (repo *UsuarioRepository) Ativa(ctx context.Context, usuarioID int) error {
	qry := `UPDATE dbo.usuario 
	          SET delet = NULL 
	        WHERE id = @p1 
			  AND delet IS NOT NULL`

	result, err := repo.db.ExecContext(ctx, qry, sql.Named("p1", usuarioID))
	if err != nil {
		return logger.LogError(fmt.Sprintf("Erro ao ativar usuario %d: %v", usuarioID, err))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.LogError(fmt.Sprintf("Erro ao verificar linhas afetadas ao ativar usuario %d: %v", usuarioID, err))
		return fmt.Errorf("erro ao verificar linhas afetadas ao ativar usuario %d: %w", usuarioID, err)
	}
	if rowsAffected == 0 {
		return logger.LogError(fmt.Sprintf("Nenhuma linha alterada para o usuario %d", usuarioID))
	}

	logger.LogError(fmt.Sprintf("Usuário ativado com sucesso: %d", usuarioID))
	return nil
}

// UsuarioExiste verifica se um usuário já existe no banco de dados.
func (repo *UsuarioRepository) UsuarioExiste(ctx context.Context, email string) (bool, error) {
	qry := `SELECT COUNT(*)
	          FROM dbo.usuario 
	         WHERE email = @p1`

	var count int
	err := repo.db.QueryRowContext(ctx, qry, sql.Named("p1", email)).Scan(&count)
	if err != nil {
		return true, logger.LogError(fmt.Sprintf("Erro ao verificar usuario %s: %v", email, err))
	}

	if count != 0 {
		return true, logger.LogError(fmt.Sprintf("Usuario já existe %s: %v", email, err))
	}
	return false, nil
}
