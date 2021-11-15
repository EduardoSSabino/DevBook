package modelos

// DadosAutenticacao cont´me o id e o token do usuário autenticado
type DadosAutenticacao struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
