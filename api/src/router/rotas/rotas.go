package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request) // campo que vai reeber a requisição e retornar a resposta
	RequerAutenticacao bool
}

// irá me retornar um router com todas as rotas configuradas

// Configurar : coloca todas as rotas dentro do router, pra isso usaremos o HandleFunc
func Configurar(r *mux.Router) *mux.Router {
	// colocando as rotas dentro do router
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...) // vai fazer um append pra cada item do slice

	for _, rota := range rotas { // iterando minhas rotas

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo) // no HandleFunc, configuramos três coisas, configuramos o URI da rota, o método da rota e também a função executada
		}
	}
	return r
}
