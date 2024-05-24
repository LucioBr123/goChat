package enum

type status_msg int
type status_usuario int

const (
	Enviado status_msg = iota + 1
	Recebido
	Vizualizado
	Excluido
	Erro
)

const (
	Disponivel status_usuario = iota + 1
	Ocupado
	Ausente
	Offline
)

// eTipoConexao

// eAcessoUsuario
