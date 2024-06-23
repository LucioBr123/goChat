package enum

type status_msg int

const (
	Enviado status_msg = iota + 1
	Recebido
	Vizualizado
	Excluido
	Erro
)

type status_usuario int

const (
	Disponivel status_usuario = iota + 1
	Ocupado
	Ausente
	Offline
)

type eTipoConexao int

const (
	Privado eTipoConexao = iota + 1
	Grupo
	Feed
)

type eAcessoUsuario int

const (
	Comum eAcessoUsuario = iota + 1
	Admin
	Ti
)
