package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

/* Middlewares : é uma camada que vai ficar entre a requisição e a resposta.
É muito utilizado qunado a gente tem alguma função que tem que ser aplicada
 pra todas as rotas, so que ao inves de você ficar entrando em rota por rota
 e colocando a função, você cria uma pacote middlewares que vai justamente fazer
 aplicação dessa função e é muito comum que no middleware você tenha uma espécie
  de alinhamento de funções */

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s ", r.Method, r.RequestURI, r.Host) // exemplo do que vai ser imprimido aqui : POST /login localhost:5000
		proximaFuncao(w, r)
	}
}

// Autenticar verifica se o usuario fazenndo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc { // HandlerFunc = func(w http.ResponseWriter, r *http.Request)
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w, r)
	}
}
