package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Rota representa todas as portas da aplicação web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotasPublicacoes...)
	rotas = append(rotas, rotaLogout)
	for _, rota := range rotas {
		if rota.RequerAutenticacao { // se for true...
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo) // alinhamento das funções
		} else { // se for false...
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer)) // Configurando o prefixo

	return router
}
