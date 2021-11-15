package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie // usado para codificar e descodificar as infromações que estão guardadas no brownser

// Configurar utiliza as variáveis de ambiente para a criação do SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar registra as informações de autenticações
func Salvar(w http.ResponseWriter, Id, token string) error {
	dados := map[string]string{
		"id":    Id,
		"token": token,
	}

	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	// colocando o cookie la no brownser
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

// Ler retornaos valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados") // lendo os cookies com os dados codificados
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)                                 // jogando um mapa vazio dentro de valores
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil { // decodificando e jogando dentro do map
		return nil, erro
	}

	return valores, nil
}

// Deletar remove os valores armazenados no cookie
func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
