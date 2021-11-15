package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router { // mux.Router é o tipo que a gente usa pra fazer a criação do router e passar as rotas pra ele depois
	r := mux.NewRouter() // essa função irá criar um router
	return rotas.Configurar(r)
}
