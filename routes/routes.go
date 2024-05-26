package routes

import (
	"github.com/LucioBr123/goChat/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cadastrarUsuario", handlers.RegisterUserHandler).Methods("POST")

	//login
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
	return r
}
