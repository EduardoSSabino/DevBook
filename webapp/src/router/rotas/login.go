package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasLogin = []Rota{
	{
		URI:                "/", // seria raiz da nossa aplicação web. OBS: temos duas rotas que fazem a mesma coisa
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login", // seria raiz da nossa aplicação web. OBS: temos duas rotas que fazem a mesma coisa
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login", // seria raiz da nossa aplicação web. OBS: temos duas rotas que fazem a mesma coisa
		Metodo:             http.MethodPost,
		Funcao:             controllers.FazerLogin,
		RequerAutenticacao: false,
	},
}
