package routes

import (
	"github.com/LucioBr123/goChat/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	// Usuario
	r.HandleFunc("/cadastrarUsuario", handlers.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/atulizaUsuario", handlers.UpdateUserHandler).Methods("POST")
	r.HandleFunc("/desativaUsuario", handlers.DeactivateUserHandler).Methods("POST")
	r.HandleFunc("/ativaUsuario", handlers.ActivateUserHandler).Methods("POST")

	//login
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
	return r
}
