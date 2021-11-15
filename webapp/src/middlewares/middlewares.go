package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger escree informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.URL.RequestURI(), r.Host)
		proximaFuncao(w, r) // passando w e r como parametro pra função
	}
}

// Autenticar verifica a existencia de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
		proximaFuncao(w, r)
	}
}
