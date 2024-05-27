package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LucioBr123/goChat/controller"
	"github.com/LucioBr123/goChat/logger"
	"github.com/LucioBr123/goChat/models"
)

// Cadastro de usuários
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	// Verifica se é um Post
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.SaveLog("Requisição inválida, Não é POST")
	}

	// Verifica se o body não tá vazio
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.SaveLog("Body vazio")
	}
	defer r.Body.Close()

	//Parseia usuario na estrutura do modelo
	var usuario models.Usuario
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.SaveLog("Erro ao decodificar JSON: " + err.Error())
		return
	}

	// Chama controller para adicionar usuario
	err = controller.CadastrarUsuario(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.SaveLog("Erro ao cadastrar usuario: " + err.Error())
		return
	}

	// Cria resposta
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}
