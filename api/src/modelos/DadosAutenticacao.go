package modelos

// DadosAutenticacao contém o toker e o id do usuário autenticado
type DadosAutenticacao struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
