package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/* Como funciona um JSON Web Token? vai pegar um série de infromações que a
 gente vai passar, como por exemplo o id do usuario e suas permissões, e vai
pegar todas essas informações e vai gerar uma string que a gente chama de token*/

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}                            // o mapa que vai ter as permissoes do token
	permissoes["authorized"] = true                          // autorização do token
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() // tempo de duração válida do meu toker
	permissoes["usuarioId"] = usuarioId
	// agora vou fazer a assinatura do token e garantir a autenticidade dele
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	// agora ou assinar o token
	return token.SignedString([]byte(config.SecretKey))

}

// ValidarToken verifica se o token passado na requisição é válido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	// Dando um Parse o meu token
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao) // o segundo parametro é a função que vai retornar pra gente a chave de verificação que  a gente vai utilizar para dar o Parse no token
	if erro != nil {
		return erro
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //chamda que valida o token
		return nil
	}
	return errors.New("Token invalido")
}

// ExtrairUsuarioId retorna o usuarioId qe está salvo no token
func ExtrairUsuarioId(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	// Dando um Parse o meu token
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao) // o segundo parametro é a função que vai retornar pra gente a chave de verificação que  a gente vai utilizar para dar o Parse no token
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64) // passo primeiro pra uma string e depois passo pra um uint64. Regras
		if erro != nil {
			return 0, erro
		}
		return usuarioId, nil
	}
	return 0, errors.New("Token invalido")
}
func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	// token tera o seguinte valor : Bearer sdfghjklçlkjhgfdsfghjklç, eu preciso somente do token e não desse "Bearer"
	if len(strings.Split(token, " ")) == 2 { // essa funçã está extraindo o token
		return strings.Split(token, " ")[1] // extraindo o token
	}
	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // se for diferente de ok
		return nil, fmt.Errorf("Metodo de assinatura inesperado %v", token.Header["alg"])
	}

	return config.SecretKey, nil // nossa chave que retorna a chave de verificação
}
