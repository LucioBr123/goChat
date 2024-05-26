package handlers

import "net/http"

//Cadastro de usuários
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	// Ensure the request method is POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	// Ensure the bdy request is not empty
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

}
