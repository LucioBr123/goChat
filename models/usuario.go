package models

type Usuario struct {
	Id             int
	Email          string
	Senha          string
	Nome           string
	Cargo          string
	FotoPath       string
	Recado         string
	EAcessoUsuario int
	EStatusUsuario int
}

func NovoUsuario(id int, email string, senha string, nome string, cargo string, fotoPath string, recado string, eAcessoUsuario int, eStatusUsuario int) *Usuario {
	return &Usuario{
		Id:             id,
		Email:          email,
		Senha:          senha,
		Nome:           nome,
		Cargo:          cargo,
		FotoPath:       fotoPath,
		Recado:         recado,
		EAcessoUsuario: eAcessoUsuario,
		EStatusUsuario: eStatusUsuario,
	}
}
