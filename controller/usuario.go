package controller

import (
	"context"

	"github.com/LucioBr123/goChat/db"
	"github.com/LucioBr123/goChat/logger"
	"github.com/LucioBr123/goChat/models"
	"github.com/LucioBr123/goChat/repository"
	"golang.org/x/crypto/bcrypt"
)

// Encrypitação de senha
func hashSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 20)
	return string(bytes), err
}

// Verifica se a senha foi correta
func checkPasswordHash(senha string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	return err == nil
}

// Cadastra um novo usuário
func CadastrarUsuario(usuario *models.Usuario) error {
	// Valida usuario
	if usuario.Nome == "" || len(usuario.Nome) < 3 || len(usuario.Nome) > 30 {
		return logger.SaveLog("Nome inválido para o : " + usuario.Nome)
	}

	if usuario.Email == "" || len(usuario.Email) < 3 || len(usuario.Email) > 50 {
		return logger.SaveLog("Email inválido para email" + usuario.Email)
	}

	if usuario.Senha == "" || len(usuario.Senha) < 4 || len(usuario.Senha) > 50 {
		return logger.SaveLog("Senha inválido para o : " + usuario.Nome)
	}

	if usuario.Cargo == "" || len(usuario.Cargo) < 2 || len(usuario.Cargo) > 30 {
		return logger.SaveLog("Cargo inválido para o : " + usuario.Cargo)
	}

	if usuario.EAcessoUsuario == 0 {
		return logger.SaveLog("Acesso inválido para o : " + usuario.Nome)
	}

	// Cria senha criptografada
	senhaCriptografada, err := hashSenha(usuario.Senha)
	if err != nil {
		return logger.SaveLog("Erro ao gerar senha criptografada: " + err.Error())
	}

	usuario.Senha = senhaCriptografada

	// Cria Repositorio para operacoes com usuario
	usuarioRepo := repository.NewUsuarioRepository(db.GetDB())

	id, err := usuarioRepo.Create(context.Background(), usuario)
	if err != nil {
		return logger.SaveLog("Erro ao criar usuario: " + err.Error())
	}

	// Atribui id ao retorno do usuario
	if id != 0 {
		usuario.Id = int(id)
	}

	return nil
}
