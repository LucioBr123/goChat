package handlers

import (
	"net/http"

	"github.com/LucioBr123/goChat/logger"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Aqui é o login"))
	logger.SaveLog("teste aqui e o login")
}
